package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const emsifaBaseURL = "https://www.emsifa.com/api-wilayah-indonesia/api"

// ProvCityService defines contract for fetching Indonesian region data
type ProvCityService interface {
	GetProvinces() (interface{}, error)
	GetCities(provID string) (interface{}, error)
	GetDetailProvince(provID string) (interface{}, error)
	GetDetailCity(cityID string) (interface{}, error)
}

type provCityService struct {
	httpClient *http.Client
}

// NewProvCityService creates a new ProvCityService instance
func NewProvCityService() ProvCityService {
	return &provCityService{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *provCityService) fetchJSON(url string) (interface{}, error) {
	resp, err := s.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external API error: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *provCityService) GetProvinces() (interface{}, error) {
	url := fmt.Sprintf("%s/provinces.json", emsifaBaseURL)
	return s.fetchJSON(url)
}

func (s *provCityService) GetCities(provID string) (interface{}, error) {
	url := fmt.Sprintf("%s/regencies/%s.json", emsifaBaseURL, provID)
	return s.fetchJSON(url)
}

func (s *provCityService) GetDetailProvince(provID string) (interface{}, error) {
	url := fmt.Sprintf("%s/province/%s.json", emsifaBaseURL, provID)
	return s.fetchJSON(url)
}

func (s *provCityService) GetDetailCity(cityID string) (interface{}, error) {
	url := fmt.Sprintf("%s/regency/%s.json", emsifaBaseURL, cityID)
	return s.fetchJSON(url)
}
