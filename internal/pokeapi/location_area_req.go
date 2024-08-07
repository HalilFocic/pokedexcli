package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint

	if pageURL != nil {
		fullUrl = *pageURL
	}
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
        return locationAreasResp, nil

	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		fmt.Print("error na new request")
		return LocationAreasResp{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Print("Error na DO")
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}
	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}
    c.cache.Add(fullUrl,data)
	return locationAreasResp, nil
}
func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseURL + endpoint

	
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
        return locationArea, nil

	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		fmt.Print("error na new request")
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Print("Error na DO")
		return LocationArea{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
    c.cache.Add(fullUrl,data)
	return locationArea, nil
}
