package main

import (
	"embed"
	"log"
	"mime"

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
	// 图片压缩用的 wasm 模块需要正确的 MIME 类型
	_ = mime.AddExtensionType(".wasm", "application/wasm")

	// Create an instance of the app structure
	// 创建一个App结构体实例
	app := NewApp()

	// Create application with options
	// 使用选项创建应用
	err := wails.Run(&options.App{
		Title:             "buddybuddy-toolkit",
		Width:             900,
		Height:            640,
		MinWidth:          900,
		MinHeight:         640,
		MaxWidth:          900,
		MaxHeight:         640,
		DisableResize:     true,
		Fullscreen:        false,
		Frameless:         true,
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
		Bind: []any{
			app,
		},
		// Windows platform specific options
		// Windows平台特定选项
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			// 让 Windows 窗口真正透明，CSS 的 12px 圆角才能完整显示，与 macOS 一致
			WindowIsTranslucent:               true,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
			BackdropType:                      windows.None,
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
			// 关闭毛玻璃材质：NSVisualEffectView 是矩形，会在 CSS 圆角外露出直角底色
			WindowIsTranslucent: false,
			About: &mac.AboutInfo{
				Title:   "buddybuddy-toolkit",
				Message: "Batch image filename toolkit (decode, sort by prefix, replace)",
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
