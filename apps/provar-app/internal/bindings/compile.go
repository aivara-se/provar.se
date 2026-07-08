package bindings

import (
	"context"
	"fmt"

	"github.com/thani-sh/provar/libs/domain"
	"github.com/thani-sh/provar/libs/engine"
	"github.com/thani-sh/provar/libs/engine/browser"
)

// Compile is the desktop app binding for the LLM compile loop. The
// editor's Compile button kicks off a stream of per-action events; the
// canvas translates those into the yellow/green/red compile borders on
// each node — that's all the editor keeps from the stream. Per-action
// Lua bodies are not persisted by this binding; the editor's generated-
// code viewer reads them later from disk via ParseFile → Action.Source.
type Compile struct {
	baseJobs
}

// Start loads the project, parses the test file, builds an LLM client
// from ~/.provar/settings.yml, opens a headless browser session, and
// launches engine.Compiler.Compile. The session is closed once the
// engine reports compile-finished (or the job is cancelled).
//
// Returning a useful error up-front matters: if the user hasn't set an
// API key yet, we don't want to open a browser for 30 seconds and then
// fail at the first LLM call. The early checks catch everything that
// doesn't need a browser session.
func (c *Compile) Start(projectDir, relPath string) (string, error) {
	project, err := domain.LoadProject(projectDir)
	if err != nil {
		return "", fmt.Errorf("load project: %w", err)
	}
	actions, err := domain.ParseFile(projectDir, relPath)
	if err != nil {
		return "", fmt.Errorf("parse test file: %w", err)
	}

	settings, err := domain.LoadSettings()
	if err != nil {
		return "", fmt.Errorf("load settings: %w", err)
	}
	if err := settings.Validate(); err != nil {
		return "", fmt.Errorf("invalid settings: %w", err)
	}
	pcfg, ok := settings.Providers[string(settings.Provider)]
	if !ok {
		return "", fmt.Errorf("no config for active provider %q", settings.Provider)
	}
	client, err := domain.ModelsClient(settings.Provider, pcfg)
	if err != nil {
		return "", fmt.Errorf("create LLM client: %w", err)
	}

	w, h := project.Browser.Resolved()
	session, err := browser.NewSession(c.Ctx, browser.Options{
		Width:  w,
		Height: h,
		// Compile always runs headless: it just walks browser tools
		// through a script the LLM authors. Showing the browser to
		// the user mid-compile would be noise.
		Headless: true,
	})
	if err != nil {
		return "", fmt.Errorf("open browser: %w", err)
	}

	ctx, cancel := context.WithCancel(c.Ctx)
	compiler := engine.NewCompiler(client)
	job, err := compiler.Compile(ctx, actions, engine.CompileOptions{
		SpecPath: relPath,
		Vars:     project.Vars,
		Browser:  session,
	})
	if err != nil {
		cancel()
		_ = session.Close()
		return "", fmt.Errorf("start compile: %w", err)
	}

	// closeSession is the cleanup that fires when the engine's
	// compile-finished event has been emitted and Job.Close() has
	// drained the Subscribe channel. Compiler doesn't own the
	// browser — caller does.
	return c.trackJob(job, cancel, func() { _ = session.Close() }), nil
}

// Cancel stops an in-flight compile by id. Same semantics as Run.Cancel.
func (c *Compile) Cancel(jobID string) error {
	return c.baseJobs.Cancel(jobID)
}
