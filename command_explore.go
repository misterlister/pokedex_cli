package main

import (
	"errors"
	"fmt"
	"strconv"
)

func commandExplore(cfg *config, locationName string) error {
	if locationName == "" {
		if cfg.currentLocation.name == nil {
			return errors.New("you must specify a location to explore")
		}
		locationName = *cfg.currentLocation.name
	}

	num, err := strconv.Atoi(locationName)

	if err == nil {
		if num >= len(cfg.currentLocationChoices) {
			return errors.New("the area number you selected is out of range")
		}
		locationName = cfg.currentLocationChoices[num]
	}

	resp, err := cfg.pokeapiClient.GetLocationArea(locationName)

	if err != nil {
		return err
	}

	cfg.currentLocation.data = &resp
	cfg.currentLocation.name = &locationName

	err = getPokemonInArea(cfg)

	if err != nil {
		return err
	}

	printPokemonInArea(cfg)

	return nil
}

func getPokemonInArea(cfg *config) error {
	if cfg.currentLocation.data == nil {
		return errors.New("you are not currently in an area")
	}
	pokemonList := make([]string, 0)
	for _, entry := range cfg.currentLocation.data.PokemonEncounters {
		pokemonList = append(pokemonList, entry.Pokemon.Name)
	}
	cfg.currentLocation.localPokemon = pokemonList
	return nil
}

func printPokemonInArea(cfg *config) {
	fmt.Printf("Pokemon in %s:\n", *cfg.currentLocation.name)
	for i, pokemon := range cfg.currentLocation.localPokemon {
		fmt.Printf("%d - %s\n", i, pokemon)
	}
}
