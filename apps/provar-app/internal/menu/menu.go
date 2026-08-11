package menu

import (
	"runtime"

	"github.com/wailsapp/wails/v2/pkg/menu"
)

// AppController abstracts the application callback methods required by native menus.
type AppController interface {
	OpenSettingsModal()
	OpenProject()
	CloseProject()
	RunActiveTest()
	CompileProject()
	ToggleTestExplorer()
}

// BuildMenu constructs and returns the native OS application menu hierarchy.
// When hasProjectOpen is false, project-scoped menus (such as View and Run) and menu items are omitted.
func BuildMenu(app AppController, hasProjectOpen bool) *menu.Menu {
	wailsMenu := menu.NewMenu()
	if runtime.GOOS == "darwin" {
		wailsMenu.Append(menu.AppMenu())
	}
	wailsMenu.Append(buildFileMenu(app, hasProjectOpen))
	if runtime.GOOS == "darwin" {
		wailsMenu.Append(menu.EditMenu())
	}
	if hasProjectOpen {
		wailsMenu.Append(buildRunMenu(app))
		wailsMenu.Append(buildViewMenu(app))
	}
	wailsMenu.Append(buildHelpMenu(app))
	return wailsMenu
}
