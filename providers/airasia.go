package providers

import (
	"errors"
	"flight-project/models"
	"flight-project/utils"
	"math/rand"
	"time"
)

type AirAsiaResponse struct {
	Status string `json:"status"`

	Flights []struct {
		FlightCode string `json:"flight_code"`

		Airline string `json:"airline"`

		FromAirport string `json:"from_airport"`

		ToAirport string `json:"to_airport"`

		DepartTime string `json:"depart_time"`

		ArriveTime string `json:"arrive_time"`

		DurationHours float64 `json:"duration_hours"`

		DirectFlight bool `json:"direct_flight"`

		Stops []struct {
			Airport string `json:"airport"`

			WaitTimeMinutes int `json:"wait_time_minutes"`
		} `json:"stops,omitempty"`

		PriceIDR float64 `json:"price_idr"`

		Seats int `json:"seats"`

		CabinClass string `json:"cabin_class"`

		BaggageNote string `json:"baggage_note"`
	} `json:"flights"`
}

func SearchAirAsia() ([]models.Flight, error) {
	delay := rand.Intn(100) + 50
	time.Sleep(time.Duration(delay) * time.Millisecond)

	err := utils.Retry(
		3,
		100*time.Millisecond,
		func() error {
			// 10% FAILURE RATE
			if rand.Intn(100) < 10 {
				return errors.New(
					"airasia provider failed",
				)
			}
			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	var response AirAsiaResponse

	err = utils.LoadJSON("data/airasia_search_response.json", &response)
	if err != nil {
		return []models.Flight{}, err
	}

	var flights []models.Flight
	for _, airasiaFlight := range response.Flights {
		flight := models.Flight{
			ID:       airasiaFlight.FlightCode + "_" + airasiaFlight.Airline,
			Provider: "AirAsia",
			Airline: models.Airline{
				Name: airasiaFlight.Airline,
				Code: airasiaFlight.FlightCode[:2],
			},
			FlightNumber: airasiaFlight.FlightCode,
			Departure: models.FlightLocation{
				Airport:  airasiaFlight.FromAirport,
				City:     "Jakarta",
				Datetime: airasiaFlight.DepartTime,
				Timestamp: utils.ParseTimestamp(
					airasiaFlight.DepartTime,
				),
			},

			Arrival: models.FlightLocation{
				Airport:  airasiaFlight.ToAirport,
				City:     "Denpasar",
				Datetime: airasiaFlight.ArriveTime,
				Timestamp: utils.ParseTimestamp(
					airasiaFlight.ArriveTime,
				),
			},

			Duration: models.FlightDuration{
				TotalMinutes: utils.ConvertHoursToMinutes(airasiaFlight.DurationHours),
				Formatted:    utils.FormatDuration(utils.ConvertHoursToMinutes(airasiaFlight.DurationHours)),
			},
			Stops: len(airasiaFlight.Stops),

			Price: models.FlightPrice{
				Amount:   airasiaFlight.PriceIDR,
				Currency: "IDR",
			},

			AvailableSeats: airasiaFlight.Seats,
			CabinClass:     airasiaFlight.CabinClass,
			Aircraft:       nil,
			Amenities:      []string{},
			Baggage: models.FlightBaggage{
				CarryOn: "Cabin baggage only",
				Checked: "Additional Fee",
			},
		}

		if !utils.IsValidFlight(flight) {
			continue
		}

		flights = append(flights, flight)
	}

	return flights, nil
}
