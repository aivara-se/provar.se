package menu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func buildViewMenu(app AppController) *menu.MenuItem {
	viewMenu := menu.NewMenu()
	viewMenu.AddText("Toggle Test Explorer", keys.CmdOrCtrl("b"), func(_ *menu.CallbackData) {
		if app != nil {
			app.ToggleTestExplorer()
		}
	})
	viewMenu.AddText("Toggle Output Console", keys.CmdOrCtrl("j"), func(_ *menu.CallbackData) {
		if app != nil {
			app.ToggleConsole()
		}
	})
	return menu.SubMenu("View", viewMenu)
}
