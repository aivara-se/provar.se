package domain

import (
	"testing"
)

func TestFileValidate_HealthyGraph(t *testing.T) {
	file := File{
		Path: ".provar/tests/login.test.yml",
		Actions: []Action{
			{ID: "open", Name: "Open Login Page", Next: []string{"enter"}},
			{ID: "enter", Name: "Enter Credentials"},
		},
	}

	report := file.Validate()
	if !report.IsValid {
		t.Errorf("expected healthy graph to be valid, got invalid with %d errors", len(report.Errors))
	}
	if len(report.Errors) != 0 {
		t.Errorf("expected 0 errors, got %d", len(report.Errors))
	}
	if len(report.Warnings) != 0 {
		t.Errorf("expected 0 warnings, got %d", len(report.Warnings))
	}
}

func TestFileValidate_EmptyGraph(t *testing.T) {
	file := File{
		Path:    ".provar/tests/empty.test.yml",
		Actions: nil,
	}

	report := file.Validate()
	if report.IsValid {
		t.Error("expected empty graph to be invalid")
	}
	if len(report.Errors) != 1 || report.Errors[0].Code != CodeEmptyGraph {
		t.Errorf("expected CodeEmptyGraph error, got %+v", report.Errors)
	}
}

func TestFileValidate_CycleDetected(t *testing.T) {
	file := File{
		Path: ".provar/tests/cycle.test.yml",
		Actions: []Action{
			{ID: "a1", Name: "Action 1", Next: []string{"a2"}},
			{ID: "a2", Name: "Action 2", Next: []string{"a1"}},
		},
	}

	report := file.Validate()
	if report.IsValid {
		t.Error("expected graph with cycle to be invalid")
	}
	hasCycle := false
	for _, err := range report.Errors {
		if err.Code == CodeCycleDetected {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		t.Errorf("expected CodeCycleDetected error, got %+v", report.Errors)
	}
}

func TestFileValidate_DanglingEdge(t *testing.T) {
	file := File{
		Path: ".provar/tests/dangling.test.yml",
		Actions: []Action{
			{ID: "a1", Name: "Action 1", Next: []string{"nonexistent"}},
		},
	}

	report := file.Validate()
	if report.IsValid {
		t.Error("expected graph with dangling edge to be invalid")
	}
	hasDangling := false
	for _, err := range report.Errors {
		if err.Code == CodeDanglingEdge && err.NodeID == "nonexistent" {
			hasDangling = true
			break
		}
	}
	if !hasDangling {
		t.Errorf("expected CodeDanglingEdge error for nonexistent, got %+v", report.Errors)
	}
}

func TestFileValidate_DisconnectedNode(t *testing.T) {
	file := File{
		Path: ".provar/tests/orphan.test.yml",
		Actions: []Action{
			{ID: "a1", Name: "Connected Action"},
			{ID: "a2", Name: "Orphaned Action"},
		},
	}

	report := file.Validate()
	if !report.IsValid {
		t.Error("warnings should not make IsValid false")
	}
	hasDisconnected := false
	for _, w := range report.Warnings {
		if w.Code == CodeDisconnectedNode && w.NodeID == "a2" {
			hasDisconnected = true
			break
		}
	}
	if !hasDisconnected {
		t.Errorf("expected CodeDisconnectedNode warning for a2, got %+v", report.Warnings)
	}
}

func TestFileValidate_MissingActionName(t *testing.T) {
	file := File{
		Path: ".provar/tests/unnamed.test.yml",
		Actions: []Action{
			{ID: "a1", Name: "   "},
		},
	}

	report := file.Validate()
	if report.IsValid {
		t.Error("expected missing action name to be invalid")
	}
	hasMissingName := false
	for _, err := range report.Errors {
		if err.Code == CodeMissingActionName && err.NodeID == "a1" {
			hasMissingName = true
			break
		}
	}
	if !hasMissingName {
		t.Errorf("expected CodeMissingActionName error, got %+v", report.Errors)
	}
}
