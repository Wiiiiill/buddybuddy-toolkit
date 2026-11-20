package main

import (
	"context"
	_ "embed"
	stdruntime "runtime"

	"github.com/energye/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed build/trayicon.png
var trayIconPNG []byte

//go:embed build/windows/icon.ico
var trayIconICO []byte

// runTray 启动系统托盘，在独立 goroutine 中运行
func (a *App) runTray() {
	systray.Run(a.onTrayReady, func() {})
}

func (a *App) onTrayReady() {
	systray.SetTooltip("buddybuddy-toolkit")
	if stdruntime.GOOS == "windows" {
		systray.SetIcon(trayIconICO)
	} else {
		systray.SetIcon(trayIconPNG)
	}

	showWindow := func() {
		runtime.WindowShow(a.ctx)
		runtime.WindowUnminimise(a.ctx)
	}

	// Windows：左键单击/双击恢复窗口，右键弹出菜单
	// macOS 保持菜单栏图标的默认行为（点击显示菜单）
	if stdruntime.GOOS == "windows" {
		systray.SetOnClick(func(menu systray.IMenu) { showWindow() })
		systray.SetOnDClick(func(menu systray.IMenu) { showWindow() })
		systray.SetOnRClick(func(menu systray.IMenu) { _ = menu.ShowMenu() })
	}

	showText, quitText := "显示主窗口", "退出"
	if a.config.Locale == "en" {
		showText, quitText = "Show window", "Quit"
	}

	showItem := systray.AddMenuItem(showText, "")
	systray.AddSeparator()
	quitItem := systray.AddMenuItem(quitText, "")

	showItem.Click(showWindow)
	quitItem.Click(func() {
		a.quitting = true
		a.stopTray()
		runtime.Quit(a.ctx)
	})
}

// stopTray 停止托盘；未初始化或重复调用时忽略错误
func (a *App) stopTray() {
	defer func() { _ = recover() }()
	systray.Quit()
}

// beforeClose 处理关闭按钮：按配置隐藏到托盘或退出，首次关闭时询问
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	if a.quitting {
		return false
	}

	switch a.config.CloseBehavior {
	case "tray":
		runtime.WindowHide(ctx)
		return true
	case "quit":
		return false
	}

	// ""（首次）或 "ask"：交给前端弹出自定义确认框
	runtime.EventsEmit(a.ctx, "close:ask")
	return true
}

// ResolveClose 处理前端关闭确认框的选择
func (a *App) ResolveClose(action string) {
	if action != "tray" && action != "quit" {
		return
	}

	// 首次选择后记住；"ask" 模式下每次都询问，不保存
	if a.config.CloseBehavior == "" {
		a.config.CloseBehavior = action
		_ = saveConfig(a.config)
	}

	if action == "tray" {
		runtime.WindowHide(a.ctx)
		return
	}

	a.quitting = true
	runtime.Quit(a.ctx)
}
