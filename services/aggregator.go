package services

import (
	"flight-project/models"
	"flight-project/providers"
	"time"
)

type ProviderResult struct {
	Flights []models.Flight
	Err     error
}

func SearchAllFlights() (models.AggregationResult, error) {
	providerFunctions := []func() ([]models.Flight, error){
		providers.SearchGaruda,
		providers.SearchLion,
		providers.SearchBatik,
		providers.SearchAirAsia,
	}

	resultsChan := make(chan ProviderResult, len(providerFunctions))

	for _, providerFunc := range providerFunctions {

		go func(searchFunc func() ([]models.Flight, error)) {
			flights, err := searchFunc()

			resultsChan <- ProviderResult{
				Flights: flights,
				Err:     err,
			}
		}(providerFunc)
	}

	var allFlights []models.Flight
	providerSucceeded := 0
	providerFailed := 0

	for range providerFunctions {
		select {
		case result := <-resultsChan:
			if result.Err != nil {
				providerFailed++
				continue
			}
			providerSucceeded++
			allFlights = append(
				allFlights,
				result.Flights...,
			)
		case <-time.After(300 * time.Millisecond):
			providerFailed++
			continue
		}
	}

	return models.AggregationResult{
		Flights:            allFlights,
		ProvidersQueried:   len(providerFunctions),
		ProvidersSucceeded: providerSucceeded,
		ProvidersFailed:    providerFailed,
	}, nil
}
