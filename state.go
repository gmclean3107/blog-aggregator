package main

import (
	"github.com/gmclean3107/blog-aggregator/internal/config"
	"github.com/gmclean3107/blog-aggregator/internal/database"
)

type State struct {
	db  *database.Queries
	cfg *config.Config
}
