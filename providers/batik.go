package providers

import (
	"flight-project/models"
	"flight-project/utils"
	"math/rand"
	"strings"
	"time"
)

type BatikResponse struct {
	Code int `json:"code"`

	Message string `json:"message"`

	Results []struct {
		FlightNumber string `json:"flightNumber"`

		AirlineName string `json:"airlineName"`

		AirlineIATA string `json:"airlineIATA"`

		Origin string `json:"origin"`

		Destination string `json:"destination"`

		DepartureDateTime string `json:"departureDateTime"`

		ArrivalDateTime string `json:"arrivalDateTime"`

		TravelTime string `json:"travelTime"`

		NumberOfStops int `json:"numberOfStops"`

		Connections []struct {
			StopAirport string `json:"stopAirport"`

			StopDuration string `json:"stopDuration"`
		} `json:"connections,omitempty"`

		Fare struct {
			BasePrice float64 `json:"basePrice"`

			Taxes float64 `json:"taxes"`

			TotalPrice float64 `json:"totalPrice"`

			CurrencyCode string `json:"currencyCode"`

			Class string `json:"class"`
		} `json:"fare"`

		SeatsAvailable int `json:"seatsAvailable"`

		AircraftModel string `json:"aircraftModel"`

		BaggageInfo string `json:"baggageInfo"`

		OnboardServices []string `json:"onboardServices"`
	} `json:"results"`
}

func SearchBatik() ([]models.Flight, error) {
	delay := rand.Intn(200) + 200
	time.Sleep(time.Duration(delay) * time.Millisecond)

	var response BatikResponse

	err := utils.LoadJSON("data/batik_air_search_response.json", &response)
	if err != nil {
		return []models.Flight{}, err
	}

	var flights []models.Flight
	for _, batikFlight := range response.Results {
		flight := models.Flight{
			ID:       batikFlight.FlightNumber + "_" + batikFlight.AirlineName,
			Provider: "Batik",
			Airline: models.Airline{
				Name: batikFlight.AirlineName,
				Code: batikFlight.AirlineIATA,
			},
			FlightNumber: batikFlight.FlightNumber,

			Departure: models.FlightLocation{
				Airport:  batikFlight.Origin,
				City:     "Jakarta",
				Datetime: batikFlight.DepartureDateTime,
				Timestamp: utils.ParseTimestamp(
					batikFlight.DepartureDateTime,
				),
			},

			Arrival: models.FlightLocation{
				Airport:  batikFlight.Destination,
				City:     "Denpasar",
				Datetime: batikFlight.ArrivalDateTime,
				Timestamp: utils.ParseTimestamp(
					batikFlight.ArrivalDateTime,
				),
			},

			Duration: models.FlightDuration{
				TotalMinutes: utils.ParseTravelTimeToMinutes(batikFlight.TravelTime),
				Formatted:    batikFlight.TravelTime,
			},

			Stops: batikFlight.NumberOfStops,

			Price: models.FlightPrice{
				Amount:   batikFlight.Fare.TotalPrice,
				Currency: batikFlight.Fare.CurrencyCode,
			},

			AvailableSeats: batikFlight.SeatsAvailable,
			CabinClass:     utils.NormalizeCabinClass(batikFlight.Fare.Class),
			Aircraft:       &batikFlight.AircraftModel,

			Amenities: batikFlight.OnboardServices,

			Baggage: models.FlightBaggage{
				CarryOn: strings.TrimSpace(strings.Split(batikFlight.BaggageInfo, ",")[0]),
				Checked: strings.TrimSpace(strings.Split(batikFlight.BaggageInfo, ",")[1]),
			},
		}

		if !utils.IsValidFlight(flight) {
			continue
		}

		flights = append(flights, flight)
	}

	return flights, nil
}
