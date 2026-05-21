package providers

import (
	"flight-project/models"
	"flight-project/utils"
	"fmt"
	"math/rand"
	"time"
)

type GarudaResponse struct {
	Status  string `json:"status"`
	Flights []struct {
		FlightID    string `json:"flight_id"`
		Airline     string `json:"airline"`
		AirlineCode string `json:"airline_code"`

		Departure struct {
			Airport  string `json:"airport"`
			City     string `json:"city"`
			Time     string `json:"time"`
			Terminal string `json:"terminal"`
		} `json:"departure"`

		Arrival struct {
			Airport  string `json:"airport"`
			City     string `json:"city"`
			Time     string `json:"time"`
			Terminal string `json:"terminal"`
		} `json:"arrival"`

		DurationMinutes int    `json:"duration_minutes"`
		Stops           int    `json:"stops"`
		Aircraft        string `json:"aircraft"`

		Price struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"price"`

		Segments []struct {
			FlightNumber string `json:"flight_number"`

			Departure struct {
				Airport string `json:"airport"`

				Time string `json:"time"`
			} `json:"departure"`

			Arrival struct {
				Airport string `json:"airport"`

				Time string `json:"time"`
			} `json:"arrival"`

			DurationMinutes int `json:"duration_minutes"`

			LayoverMinutes int `json:"layover_minutes,omitempty"`
		} `json:"segments,omitempty"`

		AvailableSeats int    `json:"available_seats"`
		FareClass      string `json:"fare_class"`

		Baggage struct {
			CarryOn int `json:"carry_on"`

			Checked int `json:"checked"`
		} `json:"baggage"`

		Amenities []string `json:"amenities,omitempty"`
	} `json:"flights"`
}

func SearchGaruda() ([]models.Flight, error) {
	delay := rand.Intn(50) + 50
	time.Sleep(time.Duration(delay) * time.Millisecond)

	var response GarudaResponse

	err := utils.LoadJSON("data/garuda_indonesia_search_response.json", &response)
	if err != nil {
		return []models.Flight{}, err
	}

	var flights []models.Flight
	for _, garudaFlight := range response.Flights {
		flight := models.Flight{
			ID:       garudaFlight.FlightID + "_" + garudaFlight.Airline,
			Provider: "Garuda",
			Airline: models.Airline{
				Name: garudaFlight.Airline,
				Code: garudaFlight.AirlineCode,
			},
			FlightNumber: garudaFlight.FlightID,

			Departure: models.FlightLocation{
				Airport:  garudaFlight.Departure.Airport,
				City:     garudaFlight.Departure.City,
				Datetime: garudaFlight.Departure.Time,
				Timestamp: utils.ParseTimestamp(
					garudaFlight.Departure.Time,
				),
			},

			Arrival: models.FlightLocation{
				Airport:  garudaFlight.Arrival.Airport,
				City:     garudaFlight.Arrival.City,
				Datetime: garudaFlight.Arrival.Time,
				Timestamp: utils.ParseTimestamp(
					garudaFlight.Arrival.Time,
				),
			},

			Duration: models.FlightDuration{
				TotalMinutes: garudaFlight.DurationMinutes,
				Formatted: utils.FormatDuration(
					garudaFlight.DurationMinutes,
				),
			},
			Stops: garudaFlight.Stops,

			Price: models.FlightPrice{
				Amount:   garudaFlight.Price.Amount,
				Currency: garudaFlight.Price.Currency,
			},

			AvailableSeats: garudaFlight.AvailableSeats,
			CabinClass:     garudaFlight.FareClass,
			Aircraft:       &garudaFlight.Aircraft,

			Amenities: garudaFlight.Amenities,

			Baggage: models.FlightBaggage{
				CarryOn: fmt.Sprintf(
					"%d kg",
					garudaFlight.Baggage.CarryOn,
				),

				Checked: fmt.Sprintf(
					"%d kg",
					garudaFlight.Baggage.Checked,
				),
			},
		}

		if !utils.IsValidFlight(flight) {
			continue
		}

		flights = append(flights, flight)
	}

	return flights, nil
}
