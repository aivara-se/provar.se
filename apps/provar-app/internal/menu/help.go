package menu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func buildHelpMenu(app AppController) *menu.MenuItem {
	helpMenu := menu.NewMenu()
	helpMenu.AddText("Provar Doctor...", nil, func(_ *menu.CallbackData) {
		if app != nil {
			app.OpenDoctorModal()
		}
	})
	helpMenu.AddSeparator()
	helpMenu.AddText("Documentation", keys.Key("f1"), func(_ *menu.CallbackData) {
		// Native documentation action handler placeholder
	})
	return menu.SubMenu("Help", helpMenu)
}
