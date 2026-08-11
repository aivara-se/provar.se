package menu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func buildFileMenu(app AppController, hasProjectOpen bool) *menu.MenuItem {
	fileMenu := menu.NewMenu()
	fileMenu.AddText("Open Project...", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
		if app != nil {
			app.OpenProject()
		}
	})
	if hasProjectOpen {
		fileMenu.AddText("Close Project", keys.CmdOrCtrl("w"), func(_ *menu.CallbackData) {
			if app != nil {
				app.CloseProject()
			}
		})
	}
	fileMenu.AddText("Settings...", keys.CmdOrCtrl(","), func(_ *menu.CallbackData) {
		if app != nil {
			app.OpenSettingsModal()
		}
	})
	return menu.SubMenu("File", fileMenu)
}
