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
	fmt.Printf("User Data:\nID: %v\nName: %s\n", user.ID, user.Name)

	err = s.cfg.SetUser(cmd.args[0])

	if err != nil {
		return err
	}

	return nil
}

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

func handlerReset(s *State, cmd Command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return err
	}
	fmt.Println("Users table reset successfully!")
	return nil
}
