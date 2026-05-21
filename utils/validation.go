package utils

import "flight-project/models"

func IsValidFlight(
	flight models.Flight,
) bool {

	if flight.Price.Amount <= 0 {
		return false
	}

	if flight.Departure.Timestamp <= 0 {
		return false
	}

	if flight.Arrival.Timestamp <= 0 {
		return false
	}

	if flight.Arrival.Timestamp <=
		flight.Departure.Timestamp {

		return false
	}

	if flight.Departure.Airport == "" {
		return false
	}

	if flight.Arrival.Airport == "" {
		return false
	}

	if flight.Duration.TotalMinutes <= 0 {
		return false
	}

	if flight.AvailableSeats < 0 {
		return false
	}

	return true
}
