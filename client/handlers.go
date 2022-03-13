package main

import (
	"os"
)

func onMenuFileConnect() {
	connectDialog.Show()
}

func onMenuFileQuit() {
	os.Exit(1)
}

func onConnectButton() {
	address, _ := connectInputAddress.GetText()
	port, _ := connectInputPort.GetText()
	connectDialog.Hide()
	go connect(address + ":" + port)
}
