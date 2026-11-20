package main

import (
	"context"
	_ "embed"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed build/trayicon.png
var trayIconPNG []byte

//go:embed build/windows/icon.ico
var trayIconICO []byte

// showMainWindow 从托盘恢复主窗口
func (a *App) showMainWindow() {
	runtime.WindowShow(a.ctx)
	runtime.WindowUnminimise(a.ctx)
}

// trayTexts 返回托盘菜单文案
func (a *App) trayTexts() (show, quit string) {
	if a.config.Locale == "en" {
		return "Show window", "Quit"
	}
	return "显示主窗口", "退出"
}

// quitFromTray 从托盘菜单退出应用
func (a *App) quitFromTray() {
	a.quitting = true
	a.stopTray()
	runtime.Quit(a.ctx)
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
