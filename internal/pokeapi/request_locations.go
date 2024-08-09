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

func (c *Client) GetLocationArea(area string) (LocationArea, error) {
	URL := BASEURL + LOCATIONURL + "/" + area

	// check if page is in the cache
	dat, ok := c.cache.Get(URL)

	if ok {
		// cache hit
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 && resp.StatusCode < 500 {
		return LocationArea{}, fmt.Errorf("error: %s is not a valid area", area)
	}

	if resp.StatusCode > 499 {
		return LocationArea{}, fmt.Errorf("error: http status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(URL, dat)

	return locationArea, nil
}
