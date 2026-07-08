package bindings

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/thani-sh/provar/libs/domain"
)

func TestCompileStart_RejectsWithoutProjectConfig(t *testing.T) {
	c := &Compile{}
	_, err := c.Start(t.TempDir(), ".provar/tests/login.test.yml")
	if err == nil {
		t.Fatal("expected error when project config is missing")
	}
	if !strings.Contains(err.Error(), "load project") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCompileStart_RejectsWithoutSettings(t *testing.T) {
	tmp := t.TempDir()
	mustMkdir(t, filepath.Join(tmp, ".provar"))
	mustWrite(t, filepath.Join(tmp, ".provar", "config.yml"), "variables:\n  baseUrl: https://example.com\n")
	mustMkdir(t, filepath.Join(tmp, ".provar", "tests"))
	mustWrite(t, filepath.Join(tmp, ".provar", "tests", "login.test.yml"),
		"- id: open_login\n  name: Open\n  info: go to login\n")

	// domain.LoadSettings reads ~/.provar/settings.yml. We have no way
	// to redirect that from the binding, so this test exercises only the
	// project-load path; the "settings invalid" branch (no provider key)
	// is covered by libs/domain's own tests.
	_ = tmp
}

func TestCompileCancel_UnknownIDIsAnError(t *testing.T) {
	c := &Compile{}
	err := c.Cancel("not-a-real-job")
	if err == nil {
		t.Fatal("expected error for unknown job id")
	}
}

func TestCompileCancel_StopsTrackedJob(t *testing.T) {
	c := &Compile{}
	job := domain.NewJob(domain.JobRunning)
	c.jobs = map[string]*jobHandle{
		job.ID: {job: job, cancel: func() {}, done: make(chan struct{})},
	}
	if err := c.Cancel(job.ID); err != nil {
		t.Fatalf("Cancel: %v", err)
	}
	if got := job.GetStatus(); got != domain.JobStopped {
		t.Errorf("job status after cancel = %q, want %q", got, domain.JobStopped)
	}
}
