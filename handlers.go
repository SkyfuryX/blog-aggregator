package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/SkyfuryX/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("A username is required")
	}
	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil || err == sql.ErrNoRows {
		return err
	}

	if err := s.config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Print("Your username has been set\n")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("A username is required")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}
	s.config.SetUser(user.Name)
	fmt.Printf("User created successfullly.\n%v\n", user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.Reset(context.Background()); err != nil {
		return err
	}
	fmt.Print("User registrations reset.\n")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == s.config.Current_user {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}
	return nil
}