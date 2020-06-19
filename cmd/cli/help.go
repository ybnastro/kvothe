package main

import (
	"fmt"
)

func showHelp() {
	fmt.Println("Welcome to Kvothe CLI!")
	fmt.Println("[PING] to execute ping, run ./kvothe-cli ping")
	fmt.Println("[Greet] to greet someone, run ./kvothe-cli greet [name]")
}
