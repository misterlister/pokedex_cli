package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationsResp, error) {
	fullURL := BASEURL + LOCATIONURL
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check if page is in the cache
	dat, ok := c.cache.Get(fullURL)

	if ok {
		// cache hit
		locationAreasResp := LocationsResp{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationsResp{}, err
		}
		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationsResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationsResp{}, fmt.Errorf("error: http status code: %v", resp.StatusCode)
	}
	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationsResp{}, err
	}

	locationAreasResp := LocationsResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationsResp{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil
}
