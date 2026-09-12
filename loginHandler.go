package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *State, cmd Command) error {

	if len(cmd.args) == 0 {
		return errors.New("Must supply username")
	}

	if len(cmd.args) > 1 {
		return errors.New("Only supply one username with login")
	}

	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err != nil {
		return err
	}

	err := s.cfg.SetUser(cmd.args[0])

	if err != nil {
		return err
	}

	fmt.Printf("Username has been set to: %s\n", cmd.args[0])

	return nil
}
