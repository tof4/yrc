package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
)

func writeToChat(message string) {
	log.Println(message)
	label, _ := gtk.LabelNew(message)
	label.SetXAlign(0)
	label.SetLineWrapMode(pango.WRAP_WORD_CHAR)
	label.SetLineWrap(true)
	ui.input.DeleteText(0, -1)
	ui.chat.Insert(label, -1)
	ui.chat.ShowAll()
	adj := ui.scrolledWindow.GetVAdjustment()
	adj.SetUpper(adj.GetUpper() + adj.GetPageSize())
	adj.SetValue(adj.GetUpper())
}
