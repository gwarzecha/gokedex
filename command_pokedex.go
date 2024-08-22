package main

import "fmt"

func callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Print("Your Pokedex is empty\n")
	} else {
		fmt.Print("Your Pokedex Entries:\n")
		for _, pokemon := range cfg.caughtPokemon {
			fmt.Printf(" - %s\n", pokemon.Name)
		}
	}

	return nil
}
