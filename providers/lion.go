package providers

import (
	"flight-project/models"
	"flight-project/utils"
	"math/rand"
	"strings"
	"time"
)

type LionResponse struct {
	Success bool `json:"success"`
	Data    struct {
		AvailableFlights []struct {
			ID string `json:"id"`

			Carrier struct {
				Name string `json:"name"`
				IATA string `json:"iata"`
			} `json:"carrier"`

			Route struct {
				From struct {
					Code string `json:"code"`
					Name string `json:"name"`
					City string `json:"city"`
				} `json:"from"`

				To struct {
					Code string `json:"code"`
					Name string `json:"name"`
					City string `json:"city"`
				} `json:"to"`
			} `json:"route"`

			Schedule struct {
				Departure         string `json:"departure"`
				DepartureTimezone string `json:"departure_timezone"`

				Arrival         string `json:"arrival"`
				ArrivalTimezone string `json:"arrival_timezone"`
			} `json:"schedule"`

			FlightTime int `json:"flight_time"`

			IsDirect bool `json:"is_direct"`

			StopCount int `json:"stop_count,omitempty"`

			Layovers []struct {
				Airport         string `json:"airport"`
				DurationMinutes int    `json:"duration_minutes"`
			} `json:"layovers,omitempty"`

			Pricing struct {
				Total    float64 `json:"total"`
				Currency string  `json:"currency"`
				FareType string  `json:"fare_type"`
			} `json:"pricing"`

			SeatsLeft int `json:"seats_left"`

			PlaneType string `json:"plane_type"`

			Services struct {
				WifiAvailable bool `json:"wifi_available"`
				MealsIncluded bool `json:"meals_included"`

				BaggageAllowance struct {
					Cabin string `json:"cabin"`
					Hold  string `json:"hold"`
				} `json:"baggage_allowance"`
			} `json:"services"`
		} `json:"available_flights"`
	} `json:"data"`
}

func SearchLion() ([]models.Flight, error) {
	delay := rand.Intn(100) + 100
	time.Sleep(time.Duration(delay) * time.Millisecond)

	var response LionResponse

	err := utils.LoadJSON("data/lion_air_search_response.json", &response)
	if err != nil {
		return []models.Flight{}, err
	}

	var flights []models.Flight
	for _, lionFlight := range response.Data.AvailableFlights {
		flight := models.Flight{
			ID:       lionFlight.ID + "_" + lionFlight.Carrier.Name,
			Provider: "Lion",
			Airline: models.Airline{
				Name: lionFlight.Carrier.Name,
				Code: lionFlight.Carrier.IATA,
			},
			FlightNumber: lionFlight.ID,

			Departure: models.FlightLocation{
				Airport:  lionFlight.Route.From.Code,
				City:     lionFlight.Route.From.City,
				Datetime: lionFlight.Schedule.Departure,
				Timestamp: utils.ParseTimestamp(
					lionFlight.Schedule.Departure,
				),
			},

			Arrival: models.FlightLocation{
				Airport:  lionFlight.Route.To.Code,
				City:     lionFlight.Route.To.City,
				Datetime: lionFlight.Schedule.Arrival,
				Timestamp: utils.ParseTimestamp(
					lionFlight.Schedule.Arrival,
				),
			},

			Duration: models.FlightDuration{
				TotalMinutes: lionFlight.FlightTime,
				Formatted: utils.FormatDuration(
					lionFlight.FlightTime,
				),
			},
			Stops: lionFlight.StopCount,

			Price: models.FlightPrice{
				Amount:   lionFlight.Pricing.Total,
				Currency: lionFlight.Pricing.Currency,
			},

			AvailableSeats: lionFlight.SeatsLeft,
			CabinClass:     strings.ToLower(lionFlight.Pricing.FareType),
			Aircraft:       &lionFlight.PlaneType,

			Amenities: []string{},

			Baggage: models.FlightBaggage{
				CarryOn: lionFlight.Services.BaggageAllowance.Cabin,
				Checked: lionFlight.Services.BaggageAllowance.Hold,
			},
		}

		if !utils.IsValidFlight(flight) {
			continue
		}

		flights = append(flights, flight)
	}

	return flights, nil
}
