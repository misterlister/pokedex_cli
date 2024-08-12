package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonData(pokemonName string) (PokemonData, error) {
	URL := BASEURL + POKEMONURL + "/" + pokemonName

	// check if page is in the cache
	dat, ok := c.cache.Get(URL)

	if ok {
		// cache hit
		pokemonData := PokemonData{}
		err := json.Unmarshal(dat, &pokemonData)
		if err != nil {
			return PokemonData{}, err
		}
		return pokemonData, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return PokemonData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 && resp.StatusCode < 500 {
		return PokemonData{}, fmt.Errorf("error: %s is not a valid area", pokemonName)
	}

	if resp.StatusCode > 499 {
		return PokemonData{}, fmt.Errorf("error: http status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonData{}, err
	}

	pokemonData := PokemonData{}
	err = json.Unmarshal(dat, &pokemonData)
	if err != nil {
		return PokemonData{}, err
	}

	c.cache.Add(URL, dat)

	return pokemonData, nil
}
