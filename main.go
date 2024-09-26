package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

// func main() {
// Example usage:
// path := "C:\\Users\\64813\\Downloads\\6668102403050861"
// prefix := "yourPrefix"
// before := "6668"
// after := "6668A"

// // Copy files based on prefix
// if err := copyFilesBasedOnPrefix(path); err != nil {
// 	fmt.Println("Error copying files:", err)
// }

// // Decode filenames
// if err := decodeFilenames(path); err != nil {
// 	fmt.Println("Error decoding filenames:", err)
// }

// // Replace filenames
//
//	if err := replaceFilenames(path, before, after); err != nil {
//		fmt.Println("Error replacing filenames:", err)
//	}
//
// }
func main() {
	// Create an instance of the app structure
	// 创建一个App结构体实例
	app := NewApp()

	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:             "buddybuddy-toolkit",
		Width:             900,
		Height:            400,
		MinWidth:          900,
		MinHeight:         400,
		MaxWidth:          1200,
		MaxHeight:         800,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         false,
		StartHidden:       false,
		HideWindowOnClose: false,
		BackgroundColour:  &options.RGBA{R: 255, G: 255, B: 255, A: 0},
		Menu:              nil,
		Logger:            nil,
		LogLevel:          logger.DEBUG,
		OnStartup:         app.startup,
		OnDomReady:        app.domReady,
		OnBeforeClose:     app.beforeClose,
		OnShutdown:        app.shutdown,
		WindowStartState:  options.Normal,
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		Bind: []interface{}{
			app,
		},
		// Windows platform specific options
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               false,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
		},
		// Mac platform specific options
		// Mac平台特定选项
		Mac: &mac.Options{
			TitleBar: mac.TitleBarDefault(),
			// TitleBar: &mac.TitleBar{
			// 	TitlebarAppearsTransparent: false,
			// 	HideTitle:                  false,
			// 	HideTitleBar:               false,
			// 	FullSizeContent:            true,
			// 	UseToolbar:                 false,
			// 	HideToolbarSeparator:       false,
			// },
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "Wails Template Vue",
				Message: "A Wails template based on Vue and Vue-Router",
				Icon:    icon,
			},
		},
		Linux: &linux.Options{
			Icon: icon,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
