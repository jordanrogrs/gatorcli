package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username not provided")
	}
	username := cmd.Args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set user: %v", err)
	}
	fmt.Printf("User has been set: %v\n", username)
	return nil
}
