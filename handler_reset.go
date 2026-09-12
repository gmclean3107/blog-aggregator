package main

import (
	"context"
	"fmt"
)

func handlerReset(s *State, cmd Command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("couldn't delete users: %v", err)
	}
	fmt.Println("Users table reset successfully!")
	return nil
}
