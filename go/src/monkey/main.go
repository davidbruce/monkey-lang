package main

import (
	"fmt"
	"os"
	"monkey/repl"
)

func main() {
	fmt.Printf("Monkey Version 0.0.1\n")
	fmt.Printf("Monkey")
	repl.Start(os.Stdin, os.Stdout)
}
