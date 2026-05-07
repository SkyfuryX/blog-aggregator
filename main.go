package main

import (
	"fmt"

	config "github.com/SkyfuryX/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	cfg.SetUser("Gannon")
	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%v\n", cfg)
}
