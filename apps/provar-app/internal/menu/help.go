package menu

import (
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func buildHelpMenu(_ AppController) *menu.MenuItem {
	helpMenu := menu.NewMenu()
	helpMenu.AddText("Documentation", keys.Key("f1"), func(_ *menu.CallbackData) {
		// Native documentation action handler placeholder
	})
	return menu.SubMenu("Help", helpMenu)
}
