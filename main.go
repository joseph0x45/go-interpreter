package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
  fmt.Printf("Hello %s! Welcome to the Monkey programming language\n", currentUser.Username)
  repl.Start(os.Stdin, os.Stdout)
}
