package bindings

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/thani-sh/provar/libs/domain"
)

func TestDoctor_RunChecks_ReturnsThreeChecks(t *testing.T) {
	withTempHome(t)
	d := Doctor{}
	checks, err := d.RunChecks("")
	if err != nil {
		t.Fatalf("RunChecks: %v", err)
	}
	if len(checks) != 3 {
		t.Fatalf("RunChecks: got %d checks, want 3", len(checks))
	}
}

func TestDoctor_RunChecks_IDs(t *testing.T) {
	withTempHome(t)
	d := Doctor{}
	checks, _ := d.RunChecks("")
	ids := make(map[string]bool, len(checks))
	for _, c := range checks {
		ids[c.ID] = true
	}
	for _, want := range []string{checkIDBrowser, checkIDAPIKey, checkIDProject} {
		if !ids[want] {
			t.Errorf("missing check ID %q in results", want)
		}
	}
}

func TestDoctor_checkProject_NoProjectOpen(t *testing.T) {
	c := checkProject("")
	if c.ID != checkIDProject {
		t.Errorf("ID = %q, want %q", c.ID, checkIDProject)
	}
	if c.Status != checkStatusWarning {
		t.Errorf("Status = %q, want %q", c.Status, checkStatusWarning)
	}
}

func TestDoctor_checkProject_MissingProvarDir(t *testing.T) {
	dir := t.TempDir()
	c := checkProject(dir)
	if c.Status != checkStatusError {
		t.Errorf("Status = %q, want %q", c.Status, checkStatusError)
	}
}

func TestDoctor_checkProject_MissingTestsDir(t *testing.T) {
	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, ".provar"), 0755); err != nil {
		t.Fatalf("mkdir .provar: %v", err)
	}
	c := checkProject(dir)
	if c.Status != checkStatusWarning {
		t.Errorf("Status = %q, want %q", c.Status, checkStatusWarning)
	}
}

func TestDoctor_checkProject_OK(t *testing.T) {
	dir := t.TempDir()
	provarDir := filepath.Join(dir, ".provar")
	if err := os.MkdirAll(filepath.Join(provarDir, "tests"), 0755); err != nil {
		t.Fatalf("mkdir .provar/tests: %v", err)
	}
	c := checkProject(dir)
	if c.Status != checkStatusOK {
		t.Errorf("Status = %q, want %q", c.Status, checkStatusOK)
	}
}

func TestDoctor_checkAPIKey_MissingSettings(t *testing.T) {
	withTempHome(t)
	c := checkAPIKey()
	if c.Status != checkStatusError {
		t.Errorf("Status = %q, want %q when settings have no API key", c.Status, checkStatusError)
	}
}

func TestDoctor_checkAPIKey_OK(t *testing.T) {
	withTempHome(t)
	s := &domain.Settings{
		Provider: domain.ProviderAnthropic,
		Providers: map[string]domain.ProviderConfig{
			string(domain.ProviderAnthropic): {
				Model:  "claude-sonnet-5",
				APIKey: "sk-test-123",
			},
		},
	}
	if err := domain.SaveSettings(s); err != nil {
		t.Fatalf("SaveSettings: %v", err)
	}
	c := checkAPIKey()
	if c.Status != checkStatusOK {
		t.Errorf("Status = %q, want %q; detail: %s", c.Status, checkStatusOK, c.Detail)
	}
}
