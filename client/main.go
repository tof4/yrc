package main

import (
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	application, _ := gtk.ApplicationNew("pl.youkai.yrc.gtk", glib.APPLICATION_FLAGS_NONE)

	application.Connect("activate", func() {
		builder, _ = gtk.BuilderNewFromFile("window.glade")
		bindUiObjects()
		bindSignals()
		application.AddWindow(ui.window)
		ui.window.Show()
	})

	os.Exit(application.Run(os.Args))
}
