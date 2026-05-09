package main

import (
	"fmt"
	"os"

	config "github.com/SkyfuryX/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}

	if len(os.Args) < 2 {
		fmt.Print("Not enough arguements provided\n")
		os.Exit(1)
	}

	s := state{&cfg}
	cmds := commands{make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	_, ok := cmds.commands[cmd.name]
	if !ok {
		fmt.Print("Invalid command\n")
		os.Exit(1)
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
