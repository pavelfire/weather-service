package wttr

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type WeatherResponse struct {
	CurrentCondition []CurrentCondition `json:"current_condition"`
}

type CurrentCondition struct {
	TempC      string `json:"temp_C"`
	FeelsLikeC string `json:"FeelsLikeC"`
}

type client struct {
	httpClient *http.Client
}

func NewClient(httpClient *http.Client) *client {
	return &client{
		httpClient: httpClient,
	}
}

func (c *client) GetTemperature(lat, long float64) (float64, error) {
	body, err := c.httpClient.Get(
		fmt.Sprintf("https://wttr.in/%f,%f?format=j1",
			lat,
			long,
		),
	)
	if err != nil {
		return 0, err
	}
	defer body.Body.Close()

	if body.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to get temperature: %s", body.Status)
	}

	var weatherResp WeatherResponse

	err = json.NewDecoder(body.Body).Decode(&weatherResp)
	if err != nil {
		return 0, err
	}
	if len(weatherResp.CurrentCondition) == 0 {
		return 0, fmt.Errorf("no weather data")
	}

	tempStr := weatherResp.CurrentCondition[0].TempC

	temp, err := strconv.ParseFloat(tempStr, 64)
	if err != nil {
		return 0, err
	}

	return temp, nil
}
