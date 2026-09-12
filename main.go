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

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerGetUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)
	commands.register("feeds", handlerGetFeeds)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	c := Command{
		command: os.Args[1],
		args:    os.Args[2:],
	}

	err = commands.run(state, c)

	if err != nil {
		log.Fatalf("error running command: %v", err)
	}

}
