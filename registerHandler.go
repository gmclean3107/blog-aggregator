package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gmclean3107/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *State, cmd Command) error {

	if len(cmd.args) == 0 {
		return errors.New("Must supply username")
	}

	if len(cmd.args) > 1 {
		return errors.New("Only supply one username with registration")
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), userParams)

	if err != nil {
		return err
	}

	fmt.Println("User created successfully")
	fmt.Printf("User Data:\n%v", user)

	err = s.cfg.SetUser(cmd.args[0])

	if err != nil {
		return err
	}

	return nil
}
