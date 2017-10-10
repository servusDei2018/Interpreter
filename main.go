package main

import (
	"fmt"
	"github.com/ash9991win/Interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the REPL \n", user.Username)
	fmt.Printf("Start typing commands! \n")
	repl.Start(os.Stdin, os.Stdout)
}
