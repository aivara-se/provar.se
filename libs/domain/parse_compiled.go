package domain

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// ParseCompiledFile reads the compiled .test.lua at the relPath of projectDir
// and returns the per-action Lua body for every action declared in it. The
// map is keyed by action ID — the same ID that appears in the matching
// .test.yml. A missing or empty compiled file returns an empty map and a
// nil error so the editor can still load the yml when the .lua hasn't been
// built yet (or when its build failed).
//
// The parse only targets the top-level `function actions.<id>(page) ... end`
// blocks the compiler emits. Nested `end` lines (Lua control flow) cannot
// appear inside the body today because translateActions only emits page-*
// statements, but the scanner matches on column-0 `end` only — so a future
// statement that happens to contain a bare `end` line is safely ignored.
func ParseCompiledFile(projectDir, relPath string) (map[string]string, error) {
	luaPath := compiledLuaPath(projectDir, relPath)
	data, err := os.ReadFile(luaPath)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]string{}, nil
		}
		return nil, err
	}
	return parseCompiledLua(string(data)), nil
}

// parseCompiledLua splits a .test.lua source string into per-action bodies.
// The format is fixed by assembleLua in libs/engine:
//
//	local actions = {}
//
//	function actions.<id>(page)
//	  <body lines>
//	end
//
//	end is recognised on its own line at column 0; deeper `end` matches are
//
// not produced by the compiler today, so a column-0 line is unambiguous.
func parseCompiledLua(src string) map[string]string {
	out := map[string]string{}
	scanner := bufio.NewScanner(strings.NewReader(src))
	// Lines emitted by the compiler are short; the scanner buffer default
	// (64KB) is plenty. If a body ever grows past that we can lift it later.
	var currentID string
	var body strings.Builder
	flush := func() {
		if currentID == "" {
			return
		}
		// Trim a single trailing newline so consumers don't see double
		// newlines when they re-emit the body verbatim.
		out[currentID] = strings.TrimRight(body.String(), "\n")
		body.Reset()
		currentID = ""
	}
	defer flush()

	for scanner.Scan() {
		line := scanner.Text()
		header, ok := luaActionHeader(line)
		if ok {
			flush()
			currentID = strings.TrimPrefix(header, "function actions.")
			currentID = strings.TrimSuffix(currentID, "(page)")
			continue
		}
		if currentID != "" && line == "end" {
			flush()
			continue
		}
		if currentID != "" {
			body.WriteString(line)
			body.WriteByte('\n')
		}
	}
	// Defensive: if the file ends mid-function (malformed), still surface
	// whatever was collected so callers aren't surprised by silence.
	flush()
	return out
}

// luaActionHeader reports whether line (with no leading whitespace) opens an
// action body, returning the full header text when it does. Examples:
//
//	"function actions.open_login_page(page)" → ("…", true)
//	"  page:navigate(...)"                    → ("", false)
func luaActionHeader(line string) (string, bool) {
	const prefix = "function actions."
	const suffix = "(page)"
	if !strings.HasPrefix(line, prefix) || !strings.HasSuffix(line, suffix) {
		return "", false
	}
	return line, true
}

// compiledLuaPath returns the .test.lua path that pairs with relPath. We
// keep the path arithmetic here (and not in project.go) because this is the
// only caller that builds the .lua side of a .test.yml pair — keeping the
// pairing logic in one file makes future renames (e.g. ".test.lua" →
// ".lua") a one-spot change.
func compiledLuaPath(projectDir, relPath string) string {
	return strings.TrimSuffix(filepath.Join(projectDir, relPath), testFileExtension) + ".test.lua"
}
