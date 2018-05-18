package main

import (
	"log"
	"os/user"
)

func getUserName() string {
	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	return user.Name
}
