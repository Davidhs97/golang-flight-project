package services

import (
	"flight-project/models"
	"math"
)

func CalculateBestValueScores(
	flights []models.Flight,
) {

	if len(flights) == 0 {
		return
	}

	maxPrice := flights[0].Price.Amount
	minPrice := flights[0].Price.Amount

	maxDuration := flights[0].Duration.TotalMinutes

	minDuration := flights[0].Duration.TotalMinutes

	maxStops := flights[0].Stops

	for _, flight := range flights {

		if flight.Price.Amount > maxPrice {
			maxPrice = flight.Price.Amount
		}

		if flight.Price.Amount < minPrice {
			minPrice = flight.Price.Amount
		}

		if flight.Duration.TotalMinutes > maxDuration {
			maxDuration = flight.Duration.TotalMinutes
		}

		if flight.Duration.TotalMinutes < minDuration {
			minDuration = flight.Duration.TotalMinutes
		}

		if flight.Stops > maxStops {
			maxStops = flight.Stops
		}
	}

	for i := range flights {

		flight := &flights[i]

		normalizedPrice := normalize(
			flight.Price.Amount,
			minPrice,
			maxPrice,
		)

		normalizedDuration := normalize(
			float64(flight.Duration.TotalMinutes),
			float64(minDuration),
			float64(maxDuration),
		)

		normalizedStops := 0.0

		if maxStops > 0 {
			normalizedStops = float64(flight.Stops) / float64(maxStops)
		}

		flight.BestValueScore = (normalizedPrice * priceWeight) + (normalizedDuration * durationWeight) + (normalizedStops * stopWeight)

		flight.BestValueScore = math.Round(flight.BestValueScore*100) / 100
	}
}

func normalize(
	value float64,
	min float64,
	max float64,
) float64 {

	if max-min == 0 {
		return 0
	}

	return (value - min) / (max - min)
}

const (
	priceWeight    = 0.5
	durationWeight = 0.3
	stopWeight     = 0.2
)
