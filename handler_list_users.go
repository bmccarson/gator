package main

import (
	"context"
	"fmt"
	"os"
)

func handlerListUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("could not get list of users")
		os.Exit(1)
	}

	currentUser := s.cfg.Current_user_name

	for _, user := range users {
		if user == currentUser {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}

	}
	return nil
}
