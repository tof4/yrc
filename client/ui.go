package main

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type uiObjects struct {
	window              *gtk.Window
	chat                *gtk.ListBox
	input               *gtk.Entry
	scrolledWindow      *gtk.ScrolledWindow
	connectDialog       *gtk.Dialog
	connectInputAddress *gtk.Entry
	connectInputPort    *gtk.Entry
}

var (
	builder *gtk.Builder
	ui      uiObjects
)

func bindUiObjects() {
	ui.window = getUiObject("main-window").(*gtk.Window)
	ui.connectDialog = getUiObject("connect-dialog").(*gtk.Dialog)
	ui.chat = getUiObject("chat").(*gtk.ListBox)
	ui.input = getUiObject("input").(*gtk.Entry)
	ui.scrolledWindow = getUiObject("scrolledWindow").(*gtk.ScrolledWindow)
	ui.connectInputAddress = getUiObject("connect-input-address").(*gtk.Entry)
	ui.connectInputPort = getUiObject("connect-input-port").(*gtk.Entry)
}

func bindSignals() {
	signals := map[string]interface{}{
		"menu-file-quit":    onMenuFileQuit,
		"menu-file-connect": onMenuFileConnect,
		"connect-button":    onConnectButton,
		"key-press-event":   onKeyPressed,
	}

	builder.ConnectSignals(signals)

	ui.window.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
		onKeyPressed(*gdk.EventKeyNewFromEvent(ev))
	})
}

func getUiObject(id string) glib.IObject {
	object, _ := builder.GetObject(id)
	return object
}
