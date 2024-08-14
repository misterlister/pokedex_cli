package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("you must specify a pokemon to inspect")
	}

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeslot := range pokemon.Types {
		fmt.Printf("\t- %s\n", typeslot.Type.Name)
	}
	return nil
}
