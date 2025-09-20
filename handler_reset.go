package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, _ command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		fmt.Println("failed to reset the users table")
		os.Exit(1)
	}
	fmt.Println("sucessfuly reset the users table")
	return nil
}
