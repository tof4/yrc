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
	chat           *gtk.ListBox
	input          *gtk.Entry
	scrolledWindow *gtk.ScrolledWindow
)

func main() {
	application, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	errorCheck(err)
	application.Connect("startup", func() {
		log.Println("application startup")
	})

	application.Connect("activate", func() {
		log.Println("application activate")

		builder, err := gtk.BuilderNewFromFile("window.glade")
		errorCheck(err)

		signals := map[string]interface{}{
			"on_main_window_destroy": onMainWindowDestroy,
		}
		builder.ConnectSignals(signals)

		winObj, err := builder.GetObject("main_window")
		chatObj, _ := builder.GetObject("chat")
		inputObj, _ := builder.GetObject("input")
		scrolledWindowObj, _ := builder.GetObject("scrolledWindow")

		win, _ := winObj.(*gtk.Window)
		chat = chatObj.(*gtk.ListBox)
		input = inputObj.(*gtk.Entry)
		scrolledWindow = scrolledWindowObj.(*gtk.ScrolledWindow)

		errorCheck(err)

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

func errorCheck(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func onMainWindowDestroy() {
	log.Println("onMainWindowDestroy")
}

func onKeyPressed(key gdk.EventKey) {
	if key.KeyVal() == gdk.KEY_Return {
		text, _ := input.GetText()
		label, _ := gtk.LabelNew(text)
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
}
