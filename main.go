package main

import (
	"fmt"
	"github.com/ruhollahh/feedly/internal/config"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", cfg)

	err = cfg.SetUser("ruhi")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", cfg)
}
