package database

import (
	"errors"

	"github.com/tof4/yrc/internal/strutil"
	"github.com/tof4/yrc/internal/validator"
)

func GetChannel(name string) (Channel, error) {
	for _, x := range channels {
		if x.Name == name {
			return x, nil
		}
	}

	return Channel{}, errors.New("Channel not found")
}

func CreateChannel(channelName string) error {
	channelName = strutil.RemoveSpaces(channelName)

	if validator.ValidateLength(channelName, 1, 20) {
		return errors.New("Invalid string length")
	}

	_, err := GetChannel(channelName)

	if err == nil {
		return errors.New("Channel already exists")
	}

	fileAppend(channelName, Paths.Channels)
	return nil
}

func AddToChannel(channelName string, username string) {

}
