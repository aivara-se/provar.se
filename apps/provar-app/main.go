package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"

	appmenu "provar-app/internal/menu"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()
	wailsMenu := appmenu.BuildMenu(app, app.ProjectOpen())

	err := wails.Run(&options.App{
		Title:  "Provar",
		Width:  1200,
		Height: 900,
		Menu:   wailsMenu,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 14, G: 17, B: 22, A: 1},
		OnStartup:        app.startup,
		Bind:             append([]interface{}{app}, app.boundBindings()...),
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHiddenInset(),
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
