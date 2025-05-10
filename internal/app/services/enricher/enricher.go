package enricher

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type Enricher interface {
	Enrich(ctx context.Context, name string) (*int, *string, *string, error)
}

type enricher struct {
}

func NewEnricher() Enricher {
	return &enricher{}
}

func (e *enricher) Enrich(ctx context.Context, FullName string) (*int, *string, *string, error) {
	encoded := url.QueryEscape(FullName)

	var age *int
	var gender, nationality *string

	type ageResp struct {
		Age int `json:"age"`
	}
	type genderResp struct {
		Gender string `json:"gender"`
	}
	type nationalityResp struct {
		Country []struct {
			CountryID string `json:"country_id"`
		} `json:"country"`
	}

	if resp, err := http.Get("https://api.agify.io/?name=" + encoded); err == nil {
		defer resp.Body.Close()
		var result ageResp
		if err := json.NewDecoder(resp.Body).Decode(&result); err == nil {
			age = &result.Age
		}
	}

	if resp, err := http.Get("https://api.genderize.io/?name=" + encoded); err == nil {
		defer resp.Body.Close()
		var result genderResp
		if err := json.NewDecoder(resp.Body).Decode(&result); err == nil {
			gender = &result.Gender
		}
	}

	if resp, err := http.Get("https://api.nationalize.io/?name=" + encoded); err == nil {
		defer resp.Body.Close()
		var result nationalityResp
		if err := json.NewDecoder(resp.Body).Decode(&result); err == nil && len(result.Country) > 0 {
			nationality = &result.Country[0].CountryID
		}
	}

	return age, gender, nationality, nil
}
