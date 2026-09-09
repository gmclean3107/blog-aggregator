package main

import (
	"log"
	"os"

	"github.com/gmclean3107/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	state := &State{&cfg}
	commands := Commands{
		commands: map[string]func(*State, Command) error{},
	}

	err = commands.register("login", handlerLogin)

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
