package main

import (
	"fmt"
	"github.com/ruhollahh/feedly/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	programState := &state{
		cfg: cfg,
	}

	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handleLogin)

	if len(os.Args) < 2 {
		log.Fatalln("command must be provided")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatalf("can't run command: %s", err)
	}
}
