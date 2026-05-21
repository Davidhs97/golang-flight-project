package services

import (
	"flight-project/models"
	"strings"
)

func SearchFlights(
	flights []models.Flight,
	request models.SearchRequest,
) []models.Flight {

	var results []models.Flight

	for _, flight := range flights {

		if strings.ToUpper(flight.Departure.Airport) !=
			strings.ToUpper(request.Origin) {
			continue
		}

		if strings.ToUpper(flight.Arrival.Airport) !=
			strings.ToUpper(request.Destination) {
			continue
		}

		if !strings.HasPrefix(
			flight.Departure.Datetime,
			request.DepartureDate,
		) {
			continue
		}

		if flight.AvailableSeats < request.Passengers {
			continue
		}

		if strings.ToLower(flight.CabinClass) !=
			strings.ToLower(request.CabinClass) {
			continue
		}

		results = append(results, flight)
	}

	return results
}
