package main

import (
	"log"
	"os"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
)

const appId = "com.github.gotk3.gotk3-examples.glade"

var (
	builder             *gtk.Builder
	chat                *gtk.ListBox
	input               *gtk.Entry
	scrolledWindow      *gtk.ScrolledWindow
	connectDialog       *gtk.Dialog
	connectInputAddress *gtk.Entry
	connectInputPort    *gtk.Entry
)

func main() {
	application, _ := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	application.Connect("activate", func() {
		log.Println("application activate")

		builder, _ = gtk.BuilderNewFromFile("window.glade")
		signals := map[string]interface{}{
			"menu-file-quit":    onMenuFileQuit,
			"menu-file-connect": onMenuFileConnect,
			"connect-button":    onConnectButton,
		}

		builder.ConnectSignals(signals)

		win := getUiObject("main-window").(*gtk.Window)
		connectDialog = getUiObject("connect-dialog").(*gtk.Dialog)
		chat = getUiObject("chat").(*gtk.ListBox)
		input = getUiObject("input").(*gtk.Entry)
		scrolledWindow = getUiObject("scrolledWindow").(*gtk.ScrolledWindow)
		connectInputAddress = getUiObject("connect-input-address").(*gtk.Entry)
		connectInputPort = getUiObject("connect-input-port").(*gtk.Entry)

		win.Show()
		application.AddWindow(win)

		win.Connect("key-press-event", func(win *gtk.Window, ev *gdk.Event) {
			onKeyPressed(*gdk.EventKeyNewFromEvent(ev))
		})
	})

	application.Connect("shutdown", func() {
		log.Println("application shutdown")
	})

	os.Exit(application.Run(os.Args))
}

func getUiObject(id string) glib.IObject {
	object, _ := builder.GetObject(id)
	return object
}

func onKeyPressed(key gdk.EventKey) {
	if key.KeyVal() == gdk.KEY_Return {
		text, _ := input.GetText()
		sendMessage(text)
		writeToChat(text)
	}
}

func writeToChat(message string) {
	label, _ := gtk.LabelNew(message)
	label.SetXAlign(0)
	label.SetLineWrapMode(pango.WRAP_WORD_CHAR)
	label.SetLineWrap(true)
	input.DeleteText(0, -1)
	chat.Insert(label, -1)
	chat.ShowAll()
	adj := scrolledWindow.GetVAdjustment()
	adj.SetUpper(adj.GetUpper() + adj.GetPageSize())
	adj.SetValue(adj.GetUpper())
}
