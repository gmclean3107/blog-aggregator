package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gmclean3107/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *State, cmd Command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.command)
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), userParams)

	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User created successfully")
	fmt.Printf("User Data:\nID: %v\nName: %s\n", user.ID, user.Name)

	err = s.cfg.SetUser(cmd.args[0])

	return nil
}

func handlerLogin(s *State, cmd Command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.command)
	}

	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err := s.cfg.SetUser(cmd.args[0])

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("Username has been set to: %s\n", cmd.args[0])

	return nil
}
