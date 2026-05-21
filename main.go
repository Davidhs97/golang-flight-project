package main

import (
	"encoding/json"
	"flight-project/models"
	"flight-project/services"
	"fmt"
	"time"
)

func main() {
	request := models.SearchRequest{
		Origin:        "CGK",
		Destination:   "DPS",
		DepartureDate: "2025-12-15",
		Passengers:    1,
		CabinClass:    "economy",
	}

	minPrice := float64(100000)
	maxPrice := float64(50000000)
	// stops := 1

	filter := models.SearchFilter{
		MinPrice: &minPrice,
		MaxPrice: &maxPrice,
		// Stops:    &stops,
	}

	start := time.Now()

	aggregationResult, err := services.SearchAllFlights()
	flights := aggregationResult.Flights

	searchTimeMs := time.Since(start).Milliseconds()

	if err != nil {
		panic(err)
	}

	metadata := models.SearchMetadata{
		TotalResults:       len(flights),
		ProvidersQueried:   aggregationResult.ProvidersQueried,
		ProvidersSucceeded: aggregationResult.ProvidersSucceeded,
		ProvidersFailed:    aggregationResult.ProvidersFailed,
		SearchTimeMs:       searchTimeMs,
		CacheHit:           false,
	}

	flights = services.SearchFlights(flights, request)

	flights = services.FilterFlights(flights, filter)

	services.CalculateBestValueScores(flights)

	services.MarkCheapestFlights(flights)

	flights = services.SortFlights(
		flights,
		models.SortByBestValue,
		models.SortAsc,
	)

	searchResponse := models.SearchResponse{
		SearchCriteria: request,
		Metadata:       metadata,
		Flights:        flights,
	}

	result, _ := json.MarshalIndent(
		searchResponse,
		"",
		"  ",
	)

	fmt.Println(string(result))
}
