package models

type SearchRequest struct {
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	DepartureDate string `json:"departure_date"`
	ReturnDate    string `json:"return_date"`
	Passengers    int    `json:"passengers"`
	CabinClass    string `json:"cabin_class"`
}

type SearchMetadata struct {
	TotalResults       int   `json:"total_results"`
	ProvidersQueried   int   `json:"providers_queried"`
	ProvidersSucceeded int   `json:"providers_succeeded"`
	ProvidersFailed    int   `json:"providers_failed"`
	SearchTimeMs       int64 `json:"search_time_ms"`
	CacheHit           bool  `json:"cache_hit"`
}

type SearchResponse struct {
	SearchCriteria SearchRequest  `json:"search_criteria"`
	Metadata       SearchMetadata `json:"metadata"`
	Flights        []Flight       `json:"flights"`
}
