package geocoding

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GeocodingResponse struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *client {
	return &client{
		httpClient: httpClient,
	}
}

func (c *client) GetCoordinates(city string) (res GeocodingResponse, err error) {
	resp, err := c.httpClient.Get(
		fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=ru&format=json",
			city))
	if err != nil {
		return GeocodingResponse{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return GeocodingResponse{}, fmt.Errorf("failed to get coordinates: %s", resp.Status)
	}

	var geoResp struct {
		Results []GeocodingResponse `json:"results"`
	}

	err = json.NewDecoder(resp.Body).Decode(&geoResp)
	if err != nil {
		return GeocodingResponse{}, err
	}

	return geoResp.Results[0], nil

}
