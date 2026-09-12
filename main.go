package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gmclean3107/blog-aggregator/internal/config"
	"github.com/gmclean3107/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.Db_url)

	if err != nil {
		log.Fatalf("error opening db connection: %v", err)
	}

	dbQueries := database.New(db)

	state := &State{dbQueries, &cfg}
	commands := Commands{
		commands: map[string]func(*State, Command) error{},
	}

	err = commands.register("login", handlerLogin)

	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}

	err = commands.register("register", handlerRegister)

	if err != nil {
		log.Fatalf("error registering command: %v", err)
	}

	args := os.Args

	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	c := Command{
		command: args[1],
		args:    args[2:],
	}

	err = commands.run(state, c)

	if err != nil {
		log.Fatalf("error running command: %v", err)
	}

}
