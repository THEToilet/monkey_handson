package main

import (
	"./repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Thie is the Monkey Programming laguage!\n", user.Username)
	fmt.Printf("Feel free to type in commnads\n")
	repl.Start(os.Stdin, os.Stdout)
}
