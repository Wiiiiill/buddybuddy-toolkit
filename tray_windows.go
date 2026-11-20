//go:build windows

package main

import "github.com/energye/systray"

func (a *App) runTray() {
	systray.Run(a.onTrayReady, func() {})
}

func (a *App) onTrayReady() {
	systray.SetTooltip("buddybuddy-toolkit")
	systray.SetIcon(trayIconICO)

	// 左键单击/双击恢复窗口，右键弹出菜单
	systray.SetOnClick(func(menu systray.IMenu) { a.showMainWindow() })
	systray.SetOnDClick(func(menu systray.IMenu) { a.showMainWindow() })
	systray.SetOnRClick(func(menu systray.IMenu) { _ = menu.ShowMenu() })

	showText, quitText := a.trayTexts()
	showItem := systray.AddMenuItem(showText, "")
	systray.AddSeparator()
	quitItem := systray.AddMenuItem(quitText, "")

	showItem.Click(a.showMainWindow)
	quitItem.Click(a.quitFromTray)
}

func (a *App) stopTray() {
	defer func() { _ = recover() }()
	systray.Quit()
}
