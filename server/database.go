package server

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type databasePaths struct {
	root     string
	users    string
	channels string
}

var paths databasePaths

func openDatabase() {
	paths.root = "ydb"
	paths.users = filepath.Join(paths.root, "usr")
	paths.channels = filepath.Join(paths.root, "chl")
	err := os.MkdirAll(paths.users, os.ModePerm)
	err = os.MkdirAll(paths.channels, os.ModePerm)
	catch(err)
}

func getUserPasswordHash(username string) (string, error) {
	usersDir, err := os.Open(paths.users)
	catch(err)
	defer usersDir.Close()

	list, _ := usersDir.Readdirnames(0)
	for _, name := range list {
		if name == username {
			passwordHash, err := os.ReadFile(filepath.Join(paths.users, name, "passwordhash"))
			catch(err)
			return strings.TrimSpace(string(passwordHash)), nil
		}
	}

	return "", errors.New("User not found")
}
