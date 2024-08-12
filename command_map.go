package main

import (
	"errors"
	"fmt"

	"github.com/misterlister/pokedex_cli/internal/pokeapi"
)

func commandMapf(cfg *config, arg string) error {

	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.locationPage++
	printMaps(resp, cfg)
	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous
	return nil
}

func commandMapb(cfg *config, arg string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you are on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.locationPage--
	printMaps(resp, cfg)
	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous
	return nil
}

func printMaps(resp pokeapi.LocationsResp, cfg *config) {
	fmt.Printf("Location areas (page %v):\n", cfg.locationPage)
	locations := make([]string, 0)
	for i, area := range resp.Results {
		fmt.Printf("%d - %s\n", i, area.Name)
		locations = append(locations, area.Name)
	}
	cfg.currentLocationChoices = locations
}
