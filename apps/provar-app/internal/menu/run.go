package menu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func buildRunMenu(app AppController) *menu.MenuItem {
	runMenu := menu.NewMenu()
	runMenu.AddText("Run Active Test", keys.CmdOrCtrl("r"), func(_ *menu.CallbackData) {
		if app != nil {
			app.RunActiveTest()
		}
	})
	runMenu.AddText("Compile Project", keys.Combo("b", keys.CmdOrCtrlKey, keys.ShiftKey), func(_ *menu.CallbackData) {
		if app != nil {
			app.CompileProject()
		}
	})
	return menu.SubMenu("Run", runMenu)
}
