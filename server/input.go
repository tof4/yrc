package server

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

func handleInput(input string, sender client) {

	validatedInput, err := validateUserInput(input)

	if err != nil {
		log.Printf("User %s sent inavlid data. Error: %s", sender.user.name, err)
		return
	}

	argumets, err := parseCommand(validatedInput)

	if len(argumets) < 1 {
		log.Printf("User %s sent inavlid data", sender.user.name)
		return
	}

	callCommand(sender, argumets)
}

func parseCommand(input string) ([]string, error) {

	r := regexp.MustCompile("'.+'|\".+\"|\\S+")
	results := r.FindAllString(input, -1)

	if len(results) < 1 {
		return results, errors.New("Too few arguments")
	}
	return results, nil
}

func validateUserInput(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("Empty input")
	}

	if len(input) > 5000 {
		return "", errors.New("Input too long")
	}

	input = strings.TrimSpace(input)

	return input, nil
}
