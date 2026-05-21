package models

type SortBy string

const (
	SortByPrice         SortBy = "price"
	SortByDuration      SortBy = "duration"
	SortByDepartureTime SortBy = "departure_time"
	SortByArrivalTime   SortBy = "arrival_time"
	SortByBestValue     SortBy = "best_value"
)

type SortOrder string

const (
	SortAsc  SortOrder = "asc"
	SortDesc SortOrder = "desc"
)
