package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/SkyfuryX/blog-aggregator/internal/database"
	"github.com/SkyfuryX/blog-aggregator/internal/rss"
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
	fmt.Printf("User created successfully.\n%v\n", user)
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

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Provide a time between requests. ex: 1s, 1m, 1h")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", cmd.args[0])
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		feed, err := s.db.GetNextFeedToFetch(context.Background(), user.ID)
		if err != nil {
			return err
		}

		err = s.db.MarkFeedFetched(context.Background(), feed.ID)
		if err != nil {
			return err
		}

		content, err := rss.FetchFeed(context.Background(), feed.Url)
		if err != nil {
			return err
		}

		fmt.Printf("%v\n", content.Channel.Title)
		fmt.Printf("%v\n", content.Channel.Description)
		fmt.Printf("%v\n\n", content.Channel.Link)
		for _, item := range content.Channel.Item {
			fmt.Printf("%v\n", item.Title)
			fmt.Printf("%v\n", item.Description)
			fmt.Printf("%v\n", item.Link)
			fmt.Printf("%v\n", item.PubDate)
			fmt.Print("\n")
		}
	}

	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("<name> and <url> are required")
	}

	feed, err := s.db.InsertFeed(context.Background(), database.InsertFeedParams{
		Name:   cmd.args[0],
		Url:    cmd.args[1],
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	result, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", result)
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %v, URL: %v, Created By: %v\n", feed.Name, feed.Url, feed.UserName)
	}
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Provide a url to follow")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	result, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Follow created.\nName: %v, User: %v\n", result.FeedName, result.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	result, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feed := range result {
		fmt.Printf("Name: %v, User: %v\n", feed.FeedName, feed.UserName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("Provide a url to unfollow")
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Follow deleted: %v", feed.Name)
	return nil
}
