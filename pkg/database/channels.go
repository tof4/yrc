package database

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/tof4/yrc/internal/errutil"
	"github.com/tof4/yrc/internal/strutil"
	"github.com/tof4/yrc/internal/validator"
	"golang.org/x/exp/slices"
)

func GetChannel(name string) (*Channel, error) {
	index := slices.IndexFunc(Channels, func(c Channel) bool { return c.Name == name })
	if index != -1 {
		return &Channels[index], nil
	}

	return &Channel{}, errors.New("Channel not found")
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

func DeleteChannel(channelName string) error {
	channelName = strutil.RemoveSpaces(channelName)
	_, err := GetChannel(channelName)

	if err != nil {
		return errors.New("Channel doesn't exists")
	}

	index := slices.IndexFunc(Channels, func(u Channel) bool { return u.Name == channelName })
	Channels = append(Channels[:index], Channels[index+1:]...)
	refreshChannelsFile()
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

	index := slices.IndexFunc(channel.Members, func(u *User) bool { return u.Name == username })
	if index != -1 {
		return errors.New("User already in channel")
	}

	channel.Members = append(channel.Members, user)
	refreshChannelsFile()
	return nil
}

func RemoveFromChannel(channelName string, username string) error {
	channel, err := GetChannel(channelName)
	if err != nil {
		return errors.New("Channel doesn't exists")
	}

	_, err = GetUser(username)
	if err != nil {
		return errors.New("User doesn't exists")
	}

	index := slices.IndexFunc(channel.Members, func(u *User) bool { return u.Name == username })
	if index == -1 {
		return errors.New("User is not a channel member")
	}

	channel.Members = append(channel.Members[:index], channel.Members[index+1:]...)
	refreshChannelsFile()
	return nil
}

func refreshChannelsFile() {
	var channelsString string
	for _, x := range Channels {
		var sb strings.Builder
		for _, y := range x.Members {
			sb.WriteString(y.Name)
			sb.WriteString(",")
		}
		channelsString += fmt.Sprintf("%s %s\n", x.Name, sb.String())
	}

	channelsFile, err := os.Create(Paths.Channels)
	defer channelsFile.Close()
	errutil.CatchFatal(err)
	channelsFile.WriteString(channelsString)
}
