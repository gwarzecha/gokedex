package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	dat, ok := c.cache.Get(url)
	if ok {
		locationAreasResp := LocationAreaResp{}
		// Unmarshal func takes data and a pointer to a struct and fill it in with the data
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreaResp{}, err
		}

		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreasResp := LocationAreaResp{}
	// Unmarshal func takes data and a pointer to a struct and fill it in with the data
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	c.cache.Add(url, dat)

	return locationAreasResp, nil
}
