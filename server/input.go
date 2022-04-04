package server

import (
	"errors"
	"log"
	"regexp"
	"strings"
)

func handleInput(input string, client yrcClient) {
	validatedInput, err := validateUserInput(input)

	if err != nil {
		log.Printf("User %s sent inavlid data. Error: %s", client.username, err)
		return
	}

	argumets, err := parseCommand(validatedInput)

	if len(argumets) < 1 {
		log.Printf("User %s sent inavlid data", client.username)
		return
	}

	callCommand(client, argumets)
}

func parseCommand(input string) ([]string, error) {
	r := regexp.MustCompile(`(\w+)||((?:\\"|[^"])*)`)
	match := r.FindAllString(input, -1)

	var results []string
	for _, s := range match {
		preparedString := strings.TrimSpace(strings.ReplaceAll(s, `\"`, ""))
		if len(preparedString) > 0 {
			results = append(results, preparedString)
		}
	}

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
