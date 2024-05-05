package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/vpaulo/wolf/bin"
	"github.com/vpaulo/wolf/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if len(os.Args) == 1 {
		fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	} else if len(os.Args) == 2 {
		bin.RunFile(os.Args[1])
	}
}
