package bindings

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/thani-sh/provar/libs/domain"
)

func TestRunStart_RejectsWithoutProjectConfig(t *testing.T) {
	r := &Run{}
	_, err := r.Start(t.TempDir(), ".provar/tests/login.test.yml", true, "")
	if err == nil {
		t.Fatal("expected error when project config is missing")
	}
	if !strings.Contains(err.Error(), "load project") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRunStart_RejectsWithoutCompiledLua(t *testing.T) {
	tmp := t.TempDir()
	mustMkdir(t, filepath.Join(tmp, ".provar"))
	mustWrite(t, filepath.Join(tmp, ".provar", "config.yml"), "variables:\n  baseUrl: https://example.com\n")
	mustMkdir(t, filepath.Join(tmp, ".provar", "tests"))
	mustWrite(t, filepath.Join(tmp, ".provar", "tests", "login.test.yml"),
		"- id: open_login\n  name: Open\n  info: go to login\n")

	r := &Run{}
	_, err := r.Start(tmp, ".provar/tests/login.test.yml", true, "")
	if err == nil {
		t.Fatal("expected error when .test.lua is missing")
	}
	if !strings.Contains(err.Error(), "compiled .test.lua") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRunStart_RejectsMalformedYml(t *testing.T) {
	tmp := t.TempDir()
	mustMkdir(t, filepath.Join(tmp, ".provar"))
	mustWrite(t, filepath.Join(tmp, ".provar", "config.yml"), "variables:\n  baseUrl: https://example.com\n")
	mustMkdir(t, filepath.Join(tmp, ".provar", "tests"))
	mustWrite(t, filepath.Join(tmp, ".provar", "tests", "broken.test.yml"),
		"key: \"unclosed\n")

	r := &Run{}
	_, err := r.Start(tmp, ".provar/tests/broken.test.yml", true, "")
	if err == nil {
		t.Fatal("expected error for malformed YAML")
	}
	if !strings.Contains(err.Error(), "parse test file") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRunCancel_UnknownIDIsAnError(t *testing.T) {
	r := &Run{}
	err := r.Cancel("not-a-real-job")
	if err == nil {
		t.Fatal("expected error for unknown job id")
	}
}

func TestRunCancel_StopsTrackedJob(t *testing.T) {
	r := &Run{}
	job := domain.NewJob(domain.JobRunning)
	r.jobs = map[string]*jobHandle{
		job.ID: {job: job, cancel: func() {}, done: make(chan struct{})},
	}
	if err := r.Cancel(job.ID); err != nil {
		t.Fatalf("Cancel: %v", err)
	}
	if got := job.GetStatus(); got != domain.JobStopped {
		t.Errorf("job status after cancel = %q, want %q", got, domain.JobStopped)
	}
}
