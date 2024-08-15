package main

import "fmt"

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to your Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
