package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/bernabe-n/pokedex/internal/pokecache"
)

func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if cache != nil {
		if val, ok := cache.Get(url); ok {
			var locationsResp RespShallowLocations
			if err := json.Unmarshal(val, &locationsResp); err == nil {
				// Cache hit
				return locationsResp, nil
			}
			// If unmarshalling fails, fall through to fetch fresh
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	if cache != nil {
		cache.Add(url, dat)
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
