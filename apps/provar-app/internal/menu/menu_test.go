package menu

import (
	"testing"
)

type mockAppController struct {
	settingsOpened      bool
	projectOpened       bool
	projectClosed       bool
	testRun             bool
	compiled            bool
	testExplorerToggled bool
}

func (m *mockAppController) OpenSettingsModal()   { m.settingsOpened = true }
func (m *mockAppController) OpenProject()         { m.projectOpened = true }
func (m *mockAppController) CloseProject()        { m.projectClosed = true }
func (m *mockAppController) RunActiveTest()       { m.testRun = true }
func (m *mockAppController) CompileProject()      { m.compiled = true }
func (m *mockAppController) ToggleTestExplorer() { m.testExplorerToggled = true }

func TestBuildMenu_NoProject(t *testing.T) {
	mock := &mockAppController{}
	m := BuildMenu(mock, false)
	if m == nil {
		t.Fatal("expected BuildMenu to return a non-nil menu")
	}

	for _, item := range m.Items {
		if item.Label == "Run" {
			t.Errorf("expected Run menu to be omitted when project is not open")
		}
		if item.Label == "View" {
			t.Errorf("expected View menu to be omitted when project is not open")
		}
	}
}

func TestBuildMenu_WithProject(t *testing.T) {
	mock := &mockAppController{}
	m := BuildMenu(mock, true)
	if m == nil {
		t.Fatal("expected BuildMenu to return a non-nil menu")
	}

	foundRun := false
	foundView := false
	for _, item := range m.Items {
		if item.Label == "Run" {
			foundRun = true
		}
		if item.Label == "View" {
			foundView = true
		}
	}

	if !foundRun {
		t.Errorf("expected Run menu to be present when project is open")
	}
	if !foundView {
		t.Errorf("expected View menu to be present when project is open")
	}
}

func TestMenuCallbacks(t *testing.T) {
	mock := &mockAppController{}
	fileItem := buildFileMenu(mock, true)
	if fileItem == nil || fileItem.SubMenu == nil {
		t.Fatal("expected buildFileMenu to return a valid submenu item")
	}

	for _, item := range fileItem.SubMenu.Items {
		if item.Click != nil {
			item.Click(nil)
		}
	}

	if !mock.projectOpened {
		t.Errorf("expected OpenProject to be called by File menu callback")
	}
	if !mock.projectClosed {
		t.Errorf("expected CloseProject to be called by File menu callback")
	}
	if !mock.settingsOpened {
		t.Errorf("expected OpenSettingsModal to be called by File menu callback")
	}

	viewItem := buildViewMenu(mock)
	if viewItem == nil || viewItem.SubMenu == nil {
		t.Fatal("expected buildViewMenu to return a valid submenu item")
	}

	for _, item := range viewItem.SubMenu.Items {
		if item.Click != nil {
			item.Click(nil)
		}
	}

	if !mock.testExplorerToggled {
		t.Errorf("expected ToggleTestExplorer to be called by View menu callback")
	}

	runItem := buildRunMenu(mock)
	if runItem == nil || runItem.SubMenu == nil {
		t.Fatal("expected buildRunMenu to return a valid submenu item")
	}

	for _, item := range runItem.SubMenu.Items {
		if item.Click != nil {
			item.Click(nil)
		}
	}

	if !mock.testRun {
		t.Errorf("expected RunActiveTest to be called by Run menu callback")
	}
	if !mock.compiled {
		t.Errorf("expected CompileProject to be called by Run menu callback")
	}
}
