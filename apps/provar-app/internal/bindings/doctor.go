package bindings

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/thani-sh/provar/libs/domain"
)

const (
	checkStatusOK      = "ok"
	checkStatusError   = "error"
	checkStatusWarning = "warning"

	checkIDBrowser = "browser"
	checkIDAPIKey  = "api_key"
	checkIDProject = "project"
)

// DoctorCheck is the result of a single environment health check.
type DoctorCheck struct {
	ID     string `json:"id"`
	Label  string `json:"label"`
	Status string `json:"status"` // "ok" | "error" | "warning"
	Detail string `json:"detail"`
}

// Doctor exposes environment diagnostic checks to the frontend.
type Doctor struct {
	BaseBinding
}

// RunChecks executes all environment health checks and returns one
// DoctorCheck per check. projectRoot may be empty when no project is
// open — the project structure check will return a warning in that case.
// Errors in individual checks are captured as DoctorCheck.Status = "error"
// so the caller always receives the full slice.
func (d Doctor) RunChecks(projectRoot string) ([]DoctorCheck, error) {
	return []DoctorCheck{
		checkBrowser(),
		checkAPIKey(),
		checkProject(projectRoot),
	}, nil
}

// checkBrowser verifies that the Playwright CLI is available on $PATH.
func checkBrowser() DoctorCheck {
	_, err := exec.LookPath("playwright")
	if err == nil {
		return DoctorCheck{
			ID:     checkIDBrowser,
			Label:  "Playwright browser binaries",
			Status: checkStatusOK,
			Detail: "playwright found on PATH",
		}
	}
	home, _ := os.UserHomeDir()
	playwrightCache := filepath.Join(home, ".cache", "ms-playwright")
	if info, statErr := os.Stat(playwrightCache); statErr == nil && info.IsDir() {
		return DoctorCheck{
			ID:     checkIDBrowser,
			Label:  "Playwright browser binaries",
			Status: checkStatusOK,
			Detail: fmt.Sprintf("browser cache found at %s", playwrightCache),
		}
	}
	return DoctorCheck{
		ID:     checkIDBrowser,
		Label:  "Playwright browser binaries",
		Status: checkStatusError,
		Detail: "playwright not found — run: npx playwright install",
	}
}

// checkAPIKey validates the active LLM provider API key from user settings.
func checkAPIKey() DoctorCheck {
	s, err := domain.LoadSettings()
	if err != nil {
		return DoctorCheck{
			ID:     checkIDAPIKey,
			Label:  "LLM API key",
			Status: checkStatusError,
			Detail: fmt.Sprintf("could not load settings: %v", err),
		}
	}
	if err := s.Validate(); err != nil {
		return DoctorCheck{
			ID:     checkIDAPIKey,
			Label:  "LLM API key",
			Status: checkStatusError,
			Detail: fmt.Sprintf("settings invalid: %v", err),
		}
	}
	return DoctorCheck{
		ID:     checkIDAPIKey,
		Label:  "LLM API key",
		Status: checkStatusOK,
		Detail: fmt.Sprintf("provider %q is configured", string(s.Provider)),
	}
}

// checkProject verifies the .provar/ directory structure under projectRoot.
func checkProject(projectRoot string) DoctorCheck {
	if projectRoot == "" {
		return DoctorCheck{
			ID:     checkIDProject,
			Label:  "Project directory structure",
			Status: checkStatusWarning,
			Detail: "no project is currently open",
		}
	}
	provarDir := filepath.Join(projectRoot, ".provar")
	if _, err := os.Stat(provarDir); err != nil {
		return DoctorCheck{
			ID:     checkIDProject,
			Label:  "Project directory structure",
			Status: checkStatusError,
			Detail: fmt.Sprintf(".provar/ directory not found in %s", projectRoot),
		}
	}
	testsDir := filepath.Join(provarDir, "tests")
	if _, err := os.Stat(testsDir); err != nil {
		return DoctorCheck{
			ID:     checkIDProject,
			Label:  "Project directory structure",
			Status: checkStatusWarning,
			Detail: ".provar/tests/ directory is missing — no tests defined yet",
		}
	}
	return DoctorCheck{
		ID:     checkIDProject,
		Label:  "Project directory structure",
		Status: checkStatusOK,
		Detail: fmt.Sprintf("project structure looks good at %s", projectRoot),
	}
}
