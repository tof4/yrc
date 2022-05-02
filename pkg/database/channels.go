package database

import "errors"

func GetChannel(name string) (Channel, error) {
	for _, x := range channels {
		if x.Name == name {
			return x, nil
		}
	}

	return Channel{}, errors.New("Channel not found")
}
