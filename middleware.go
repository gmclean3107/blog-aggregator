package main

import (
	"context"

	"github.com/gmclean3107/blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {

		username := s.cfg.Current_user_name

		user, err := s.db.GetUser(context.Background(), username)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
