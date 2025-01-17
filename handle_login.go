package main

import "fmt"

func handleLogin(state *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("not enough or too many arguments passed")
	}

	userName := cmd.args[0]

	err := state.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("setUser: %w", err)
	}

	fmt.Println("User switched successfully!")

	return nil
}
