package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Couldn't set user: %w", err)
	}
	fmt.Printf("User has been set to %s", s.cfg.Current_user_name)
	return nil
}
