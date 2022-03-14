package client

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type uiObjects struct {
	window         *gtk.Window
	chat           *gtk.ListBox
	input          *gtk.Entry
	scrolledWindow *gtk.ScrolledWindow

	menuConnectButton    *gtk.MenuItem
	menuDisconnectButton *gtk.MenuItem

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
	ui.chat = getUiObject("chat").(*gtk.ListBox)
	ui.input = getUiObject("input").(*gtk.Entry)
	ui.scrolledWindow = getUiObject("scrolled-window").(*gtk.ScrolledWindow)

	ui.menuConnectButton = getUiObject("menu-connect-button").(*gtk.MenuItem)
	ui.menuDisconnectButton = getUiObject("menu-disconnect-button").(*gtk.MenuItem)

	ui.connectDialog = getUiObject("connect-dialog").(*gtk.Dialog)
	ui.connectInputAddress = getUiObject("connect-input-address").(*gtk.Entry)
	ui.connectInputPort = getUiObject("connect-input-port").(*gtk.Entry)
}

func bindSignals() {
	signals := map[string]interface{}{
		"menu-quit":       onMenuQuit,
		"menu-connect":    onMenuConnect,
		"menu-disconnect": onMenuDisconnect,
		"connect-button":  onConnectDialogButton,
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

func switchConnectButton() {
	glib.IdleAdd(func() {
		if ui.menuConnectButton.GetVisible() {
			ui.menuConnectButton.Hide()
			ui.menuDisconnectButton.Show()
		} else {
			ui.menuConnectButton.Show()
			ui.menuDisconnectButton.Hide()
		}
	})
}
