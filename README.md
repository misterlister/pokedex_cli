# Pokedex CLI

**Pokedex CLI** is a command line application written in Go that lets you explore Pokémon regions, catch Pokémon, and inspect the stats of the Pokémon you've caught. It uses the [PokeAPI](https://pokeapi.co/) to retrieve up-to-date Pokémon data and provides an interactive REPL-style interface to navigate between locations and manage your Pokémon collection.

## Features

- **Explore Pokémon regions:** Navigate through different locations and areas to discover Pokémon native to those spots.
- **Catch Pokémon:** Attempt to catch Pokémon found in the current location to add them to your personal Pokedex.
- **Inspect Pokémon:** View detailed stats for caught Pokémon, including height, weight, types, and base stats.
- **Manage Pokedex:** View a list of all Pokémon you've caught so far.
- **Pagination:** Navigate forwards and backwards through location pages to explore the full Pokémon world.

## Installation

1. Ensure you have [Go](https://golang.org/dl/) installed (version 1.16+ recommended).  
2. Clone the repository:  
   ```bash
   git clone https://github.com/misterlister/pokedex_cli
   ```
3. Build the application:
    ```bash
    cd pokedex_cli
    go build
    ```
4. Run the application
    ```bash
    ./pokedex_cli
    ```

## Usage

After launching the application, you will be greeted with a prompt:

``` bash
Pokedex >
```
You can type commands to interact with the app. Some useful commands include:

| Command              | Description                                                                     |
|----------------------|---------------------------------------------------------------------------------|
| `help`               | Show all valid commands and their descriptions                                  |
| `map`                | Show the next page of available map locations                                   |
| `mapb`               | Show the previous page of map locations                                         |
| `explore <location>` | List all Pokémon found in the specified location (e.g. `explore mt-coronet-2f`) |
| `catch <pokemon>`    | Attempt to catch a specified Pokémon in the current area (e.g. `catch pikachu`) |
| `inspect <pokemon>`  | Display details of a caught Pokémon including stats (e.g. `inspect pikachu`)    |
| `pokedex`            | List all Pokémon you have caught so far                                         |
| `exit`               | Exit the application                                                            |

## Dependencies

- [PokeAPI](https://pokeapi.co/) — RESTful Pokémon data API  
- Internal packages:  
  `github.com/misterlister/pokedex_cli/internal/pokeapi` — PokeAPI client  
  `github.com/misterlister/pokedex_cli/internal/pokecache` — Caching layer for API data


## License

MIT License