package client

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
)

func logWriteMessage(message string) {
	label, _ := gtk.LabelNew(message)
	logWriteToChat(label)
}

func logWriteError(message string) {
	label, _ := gtk.LabelNew("")
	label.SetMarkup("<span color=\"red\" style=\"italic\">" + message + "</span>")
	logWriteToChat(label)
}

func logWriteStatus(message string) {
	label, _ := gtk.LabelNew("")
	label.SetMarkup("<span color=\"blue\" style=\"italic\">" + message + "</span>")
	logWriteToChat(label)
}

func logWriteToChat(label *gtk.Label) {
	label.SetXAlign(0)
	label.SetLineWrapMode(pango.WRAP_WORD_CHAR)
	label.SetLineWrap(true)
	label.SetSelectable(true)

	glib.IdleAdd(func() {
		ui.input.DeleteText(0, -1)
		ui.chat.Insert(label, -1)
		ui.chat.ShowAll()
		adj := ui.scrolledWindow.GetVAdjustment()
		adj.SetUpper(adj.GetUpper() + adj.GetPageSize())
		adj.SetValue(adj.GetUpper())
	})
}
