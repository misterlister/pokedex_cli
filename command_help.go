package main

import "fmt"

func commandHelp(cfg *config, arg string) error {
	fmt.Println(DIVIDER)
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Valid commands:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println(DIVIDER)
	return nil
}
