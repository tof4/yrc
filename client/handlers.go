package main

import (
	"os"

	"github.com/gotk3/gotk3/gdk"
)

func onMenuFileConnect() {
	ui.connectDialog.Show()
}

func onMenuFileQuit() {
	os.Exit(1)
}

func onConnectButton() {
	address, _ := ui.connectInputAddress.GetText()
	port, _ := ui.connectInputPort.GetText()
	ui.connectDialog.Hide()
	go connect(address + ":" + port)
}

func onKeyPressed(key gdk.EventKey) {
	if key.KeyVal() == gdk.KEY_Return {
		text, _ := ui.input.GetText()
		sendMessage(text)
		logWriteMessage("me: " + text)
	}
}
