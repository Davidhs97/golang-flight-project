package services

import (
	"flight-project/models"
	"fmt"
)

func MarkCheapestFlights(
	flights []models.Flight,
) {

	cheapestMap := make(map[string]int)

	for i, flight := range flights {

		groupKey := fmt.Sprintf(
			"%s-%s-%d",
			flight.Departure.Airport,
			flight.Arrival.Airport,
			flight.Departure.Timestamp,
		)

		existingIndex, exists := cheapestMap[groupKey]

		if !exists {
			cheapestMap[groupKey] = i

			continue
		}

		if flight.Price.Amount < flights[existingIndex].Price.Amount {

			cheapestMap[groupKey] = i
		}
	}

	for _, cheapestIndex := range cheapestMap {

		flights[cheapestIndex].IsCheapest = true
	}
}
