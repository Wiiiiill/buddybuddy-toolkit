//go:build darwin

package main

import "github.com/getlantern/systray"

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

	go func() {
		for {
			select {
			case <-showItem.ClickedCh:
				a.showMainWindow()
			case <-quitItem.ClickedCh:
				a.quitFromTray()
				return
			}
		}
	}()
}

func (a *App) stopTray() {
	defer func() { _ = recover() }()
	systray.Quit()
}
