package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, handler func(*state, command) error) {
	c.handlers[name] = handler
}

func (c *commands) run(state *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("command not found")
	}

	err := handler(state, cmd)
	if err != nil {
		return fmt.Errorf("handler: %w", err)
	}

	return nil
}
