package main

import (
	"context"
	"errors"
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

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		fmt.Printf("usage: %s <feed name> <url>\n", cmd.Name)
		os.Exit(1)
	}

	usr, err := s.db.GetUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("could not get %s from database", s.cfg.Current_user_name)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    usr.ID,
	})
	fmt.Printf("Name: %s\nURL: %s", feed.Name, feed.Url)
	return nil
}

func handlerListFeeds(s *state, _ command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return errors.New("could not get list of feeds")
	}

	for _, feed := range feeds {
		usrName, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("could not get user for id %s", feed.UserID)
		}
		fmt.Printf("Feed Name: %s\nURL: %s\nUser Name: %s\n", feed.Name, feed.Url, usrName)
		fmt.Println("-----------------------")
	}
	return nil
}
