package models

type AggregationResult struct {
	Flights            []Flight
	ProvidersQueried   int
	ProvidersSucceeded int
	ProvidersFailed    int
}
