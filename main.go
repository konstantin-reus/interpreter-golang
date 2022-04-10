package main

import (
	"fmt"
	"os"
	"os/user"

	"com.reus.intepreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is ololo language!\n", user.Username)
	fmt.Printf("Enter commands")
	repl.Start(os.Stdin, os.Stdout)
}
