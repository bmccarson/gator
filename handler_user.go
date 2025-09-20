package main

import (
	"context"
	"fmt"
	"os"
)

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
