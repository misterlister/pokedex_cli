package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, arg string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you haven't caught any pokemon yet")
	}
	fmt.Println("Your pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("\t-%s\n", pokemon.Name)
	}
	return nil
}
