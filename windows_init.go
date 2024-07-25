//go:build windows
// +build windows

package main

import (
	"github.com/lxn/win"
	"syscall"
)

func InitWithApp() {
	wn, _ := syscall.UTF16PtrFromString("wailsdemo")
	hwnd := win.FindWindow(nil, wn)
	win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)
}
