package main

import (
	"errors"
	"fmt"

	"github.com/misterlister/pokedex_cli/internal/pokeapi"
)

func commandExplore(cfg *config, locationName string) error {
	if locationName == "" {
		return errors.New("you must specify a location to explore")
	}

	resp, err := cfg.pokeapiClient.GetLocationArea(locationName)

	if err != nil {
		return err
	}

	printPokemonInArea(resp, locationName)

	return nil
}

func printPokemonInArea(resp pokeapi.LocationArea, locationName string) {
	fmt.Printf("Pokemon in %s:\n", locationName)
	for _, entry := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", entry.Pokemon.Name)
	}
}
