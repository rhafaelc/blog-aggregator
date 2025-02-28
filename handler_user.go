package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rhafaelc/blog-aggregator/internal/database"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage %v", cmd.Name)
	}
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get users: %w", err)
	}
	for i := 0; i < len(users); i++ {
		if users[i].Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", users[i].Name)
		} else {
			fmt.Printf("* %s\n", users[i].Name)
		}
	}
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	username := cmd.Arguments[0]
	timeNow := time.Now().UTC()
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Name:      username,
	}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("user created: \n")
	printUser(user)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Arguments[0]
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Printf("user switched: %v\n", s.cfg.CurrentUserName)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
