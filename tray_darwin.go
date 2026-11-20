//go:build darwin

package main

/*
#cgo darwin CFLAGS: -x objective-c -fobjc-arc
#cgo darwin LDFLAGS: -framework Cocoa
#include <stdlib.h>

void tray_setup(const char *showText, const char *quitText, const char *tooltip, const void *iconData, int iconLen);
void tray_remove(void);
*/
import "C"

import "unsafe"

var trayAppRef *App

// runTray 在 Wails 已有的主线程 NSApplication 上挂一个状态栏图标，
// 不启动自己的原生循环（systray 库会因此与 Wails 的主线程冲突）。
func (a *App) runTray() {
	trayAppRef = a
	show, quit := a.trayTexts()

	cShow := C.CString(show)
	cQuit := C.CString(quit)
	cTip := C.CString("buddybuddy-toolkit")
	defer C.free(unsafe.Pointer(cShow))
	defer C.free(unsafe.Pointer(cQuit))
	defer C.free(unsafe.Pointer(cTip))

	var icon unsafe.Pointer
	if len(trayIconPNG) > 0 {
		icon = unsafe.Pointer(&trayIconPNG[0])
	}
	C.tray_setup(cShow, cQuit, cTip, icon, C.int(len(trayIconPNG)))
}

func (a *App) stopTray() {
	C.tray_remove()
}

//export tray_on_show
func tray_on_show() {
	if trayAppRef != nil {
		go trayAppRef.showMainWindow()
	}
}

//export tray_on_quit
func tray_on_quit() {
	if trayAppRef != nil {
		go trayAppRef.quitFromTray()
	}
}
