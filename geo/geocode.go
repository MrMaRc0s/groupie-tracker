package geo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// curl -H "X-Api-Key: DUuHzfSg6h9fVT+twOX7Kg==61bhJkCixkC6QN8V" "https://api.api-ninjas.com/v1/geocoding?city=Manchester&country=uk"
// REPLACE UK WITH GB
// REPLACE USA WITH US
// REPLACE PORTUGAL WITH PT
func Geocode(city, country string) (float64, float64, error) {
	apiKey := "DUuHzfSg6h9fVT+twOX7Kg==61bhJkCixkC6QN8V"
	if apiKey == "" {
		return 0, 0, fmt.Errorf("API key missing")
	}

	baseURL := "https://api.api-ninjas.com/v1/geocoding"
	params := url.Values{}
	params.Add("city", city)
	params.Add("country", country)
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, _ := http.NewRequest("GET", fullURL, nil)
	req.Header.Set("X-Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("API error: %s", resp.Status)
	}

	var results []struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return 0, 0, err
	}

	if len(results) == 0 {
		return 0, 0, fmt.Errorf("no coordinates found")
	}

	return results[0].Latitude, results[0].Longitude, nil
}
