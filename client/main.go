package client

import (
	"os"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func Initialize() {
	application, _ := gtk.ApplicationNew("pl.youkai.yrc.gtk", glib.APPLICATION_NON_UNIQUE)

	application.Connect("activate", func() {
		builder, _ = gtk.BuilderNewFromFile("../../client/window.glade")
		bindUiObjects()
		bindSignals()
		application.AddWindow(ui.window)
		ui.window.Show()
	})

	os.Exit(application.Run(os.Args))
}
