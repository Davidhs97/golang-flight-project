package services

import "flight-project/models"

func FilterFlights(
	flights []models.Flight,
	filter models.SearchFilter,
) []models.Flight {

	var filtered []models.Flight

	for _, flight := range flights {

		if filter.MinPrice != nil &&
			flight.Price.Amount < *filter.MinPrice {

			continue
		}

		if filter.MaxPrice != nil &&
			flight.Price.Amount > *filter.MaxPrice {

			continue
		}

		if filter.Stops != nil &&
			flight.Stops != *filter.Stops {

			continue
		}

		if filter.Airline != nil &&
			flight.Airline.Name != *filter.Airline {

			continue
		}

		if filter.MaxDuration != nil &&
			flight.Duration.TotalMinutes > *filter.MaxDuration {

			continue
		}

		filtered = append(filtered, flight)
	}

	return filtered
}
