//go:build linux

package main

import "github.com/energye/systray"

func (a *App) runTray() {
	systray.Run(a.onTrayReady, func() {})
}

func (a *App) onTrayReady() {
	systray.SetTooltip("buddybuddy-toolkit")
	systray.SetIcon(trayIconPNG)

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
