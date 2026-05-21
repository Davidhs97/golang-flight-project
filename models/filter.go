package models

type SearchFilter struct {
	MinPrice      *float64
	MaxPrice      *float64
	Stops         *int
	DepartureTime *string
	ArrivalTime   *string
	Airline       *string
	MaxDuration   *int
}
