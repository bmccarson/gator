package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bmccarson/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		fmt.Printf("user %s already exsist\n", user.Name)
		os.Exit(1)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Couldn't set user: %w\n", err)
	}

	fmt.Printf("user %s has been logged in\n", user.Name)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		fmt.Printf("user %s does not exsist\n", user.Name)
		os.Exit(1)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Couldn't set user: %w\n", err)
	}
	fmt.Printf("User has been set to %s\n", s.cfg.Current_user_name)
	return nil
}

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

func handlerReset(s *state, _ command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		fmt.Println("failed to reset the users table")
		os.Exit(1)
	}
	fmt.Println("sucessfuly reset the users table")
	return nil
}

func handlerAgg(s *state, _ command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
