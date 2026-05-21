package models

type Flight struct {
	ID             string         `json:"id"`
	Provider       string         `json:"provider"`
	Airline        Airline        `json:"airline"`
	FlightNumber   string         `json:"flight_number"`
	Departure      FlightLocation `json:"departure"`
	Arrival        FlightLocation `json:"arrival"`
	Duration       FlightDuration `json:"duration"`
	Stops          int            `json:"stops"`
	Price          FlightPrice    `json:"price"`
	AvailableSeats int            `json:"available_seats"`
	CabinClass     string         `json:"cabin_class"`
	Aircraft       *string        `json:"aircraft"`
	Amenities      []string       `json:"amenities"`
	Baggage        FlightBaggage  `json:"baggage"`
	BestValueScore float64        `json:"-"`
	IsCheapest     bool           `json:"-"`
}

type Airline struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type FlightLocation struct {
	Airport   string `json:"airport"`
	City      string `json:"city"`
	Datetime  string `json:"datetime"`
	Timestamp int64  `json:"timestamp"`
}

type FlightDuration struct {
	TotalMinutes int    `json:"total_minutes"`
	Formatted    string `json:"formatted"`
}

type FlightPrice struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type FlightBaggage struct {
	CarryOn string `json:"carry_on"`
	Checked string `json:"checked"`
}
