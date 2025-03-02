package main 

import "errors"

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if handler, ok := c.registeredCommands[cmd.Name]; ok {
		return handler(s, cmd)
	} else {
		return errors.New("command not found")
	}
}

