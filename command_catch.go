package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

func commandCatch(cfg *config, pokemonName string) error {
	if cfg.currentLocation.name == nil {
		return errors.New("you must explore an area to catch a pokemon")
	}
	if pokemonName == "" {
		return errors.New("you must specify a pokemon to capture")
	}

	num, err := strconv.Atoi(pokemonName)

	if err == nil {
		if num >= len(cfg.currentLocation.localPokemon) {
			return errors.New("the pokemon number you selected is out of range")
		}
		pokemonName = cfg.currentLocation.localPokemon[num]
	}

	pokemonInArea := false
	for _, name := range cfg.currentLocation.localPokemon {
		if name == pokemonName {
			pokemonInArea = true
			break
		}
	}

	if !pokemonInArea {
		return errors.New("the pokemon you selected is not in the current area")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemonData(pokemonName)

	if err != nil {
		return err
	}

	if catchAttempt(pokemonName, pokemon.BaseExperience) {
		cfg.caughtPokemon[pokemonName] = pokemon
	}

	return nil
}

func catchAttempt(pokemonName string, catchChance int) bool {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	attemptVal := rand.Intn(catchChance)
	if attemptVal < CATCH_THRESHOLD {
		fmt.Printf("%s was caught!\n", pokemonName)
		return true
	}
	fmt.Printf("%s escaped!\n", pokemonName)
	return false
}
