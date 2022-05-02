package database

import (
	"errors"
	"path/filepath"
)

func SaveMessage(channelName string, message string) {
	path := filepath.Join(Paths.Chat, channelName)
	fileAppend(message, path)
}

func GetChannelMessages(channelName string, amount int) ([]string, error) {
	_, err := GetChannel(channelName)

	if err != nil {
		return []string{}, err
	}

	if amount < 1 || amount > 1000 {
		return []string{}, errors.New("Invalid amount")
	}
	amount++

	path := filepath.Join(Paths.Chat, channelName)
	return BackwardFileRead(path, amount), nil
}
