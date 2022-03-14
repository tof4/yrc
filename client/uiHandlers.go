package main

import (
	"os"

	"github.com/gotk3/gotk3/gdk"
)

func onMenuConnect() {
	ui.connectDialog.Show()
}

func onMenuDisconnect() {
	connection.Close()
}

func onMenuQuit() {
	os.Exit(1)
}

func onConnectDialogButton() {
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
