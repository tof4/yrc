package database

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/tof4/yrc/internal/errutil"
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

func AddToChannel(channelName string, username string) error {
	channel, err := GetChannel(channelName)

	if err != nil {
		return errors.New("Channel doesn't exists")
	}

	user, err := GetUser(username)

	if err != nil {
		return errors.New("User doesn't exists")
	}

	for _, x := range channel.Members {
		if x.Name == username {
			return errors.New("User already in channel")
		}
	}

	channel.Members = append(channel.Members, &user)

	channelsFile, err := os.OpenFile(Paths.Channels, os.O_WRONLY, 0600)
	defer channelsFile.Close()
	errutil.CatchFatal(err)

	var channelsString string
	for _, x := range channels {
		var userStrings []string
		for _, y := range channel.Members {
			userStrings = append(userStrings, y.Name)
		}
		channelsString += fmt.Sprintf("%s %s,\n", x.Name, strings.Join(userStrings, ","))
	}

	channelsFile.WriteString(channelsString)
	return nil
}
