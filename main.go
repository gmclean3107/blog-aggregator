package main

import (
	"fmt"
	"log"

	"github.com/gmclean3107/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	err = cfg.SetUser("Gerard")

	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err = config.Read()

	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Println(cfg)
}
