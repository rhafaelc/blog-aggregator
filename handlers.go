package main 

import (
	"errors"
	"fmt"

)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	username := cmd.Arguments[0]
	s.Config.SetUser(username)
	fmt.Printf("username has been set: %v\n", s.Config.CurrentUserName)
	return nil
}


