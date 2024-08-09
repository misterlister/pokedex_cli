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
	printMaps(resp, cfg.locationPage)
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
	printMaps(resp, cfg.locationPage)
	cfg.nextLocationsURL = resp.Next
	cfg.prevLocationsURL = resp.Previous
	return nil
}

func printMaps(resp pokeapi.LocationsResp, page int) {
	fmt.Printf("Location areas (page %v):\n", page)

	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}
}
