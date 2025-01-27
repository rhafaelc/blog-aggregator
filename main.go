package main

import (
	"fmt"
	"log"

	"github.com/rhafaelc/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	cfg.SetUser("rhafaelc")
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Println(cfg)
}
