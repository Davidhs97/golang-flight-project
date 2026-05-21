package services

import (
	"flight-project/models"
	"sort"
)

func SortFlights(
	flights []models.Flight,
	sortBy models.SortBy,
	order models.SortOrder,
) []models.Flight {

	sort.Slice(flights, func(i, j int) bool {
		switch sortBy {
		case models.SortByPrice:
			if order == models.SortAsc {
				return flights[i].Price.Amount <
					flights[j].Price.Amount
			}
			return flights[i].Price.Amount >
				flights[j].Price.Amount

		case models.SortByDuration:
			if order == models.SortAsc {
				return flights[i].Duration.TotalMinutes <
					flights[j].Duration.TotalMinutes
			}
			return flights[i].Duration.TotalMinutes >
				flights[j].Duration.TotalMinutes

		case models.SortByDepartureTime:
			if order == models.SortAsc {
				return flights[i].Departure.Datetime <
					flights[j].Departure.Datetime
			}
			return flights[i].Departure.Datetime >
				flights[j].Departure.Datetime

		case models.SortByArrivalTime:
			if order == models.SortAsc {
				return flights[i].Arrival.Datetime <
					flights[j].Arrival.Datetime
			}
			return flights[i].Arrival.Datetime >
				flights[j].Arrival.Datetime

		case models.SortByBestValue:
			if flights[i].BestValueScore ==
				flights[j].BestValueScore {

				return flights[i].IsCheapest
			}

			if order == models.SortAsc {

				return flights[i].BestValueScore <
					flights[j].BestValueScore
			}

			return flights[i].BestValueScore >
				flights[j].BestValueScore
		}

		return true
	})

	return flights
}
