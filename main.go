package main

import (
	"fmt"
	"log"

	"github.com/jordanrogrs/gatorcli/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
		return
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("jordan")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
		return
	}
	fmt.Println(cfg)

}
