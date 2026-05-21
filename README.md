# Flight Search & Aggregation System

A Golang-based flight aggregation system that fetches mock flight data from multiple airline providers, normalizes the responses into a unified format, and supports searching, filtering, sorting, concurrency, retry logic, timeout handling, validation, and best-value ranking.

Built as part of the Software Engineer Take-Home Assignment.

---

# Features

## Flight Aggregation

Aggregates flight data from multiple airline providers:

* Garuda Indonesia
* Lion Air
* Batik Air
* AirAsia

All provider responses are normalized into a unified internal flight model.

---

## Data Normalization

Handles provider data inconsistencies such as:

* Different datetime formats
* Different duration formats
* Different cabin class formats
* Different baggage structures
* Different pricing schemas

Example:

* `1h 45m` → normalized to total minutes
* `Y` → normalized to `economy`

---

## Search

Supports searching by:

* Origin
* Destination
* Departure date

---

## Filtering

Supports filtering by:

* Minimum price
* Maximum price
* Number of stops
* Airlines
* Duration

---

## Sorting

Supports sorting by:

* Price
* Duration
* Departure time
* Arrival time
* Best value score

Supports:

* Ascending order
* Descending order

---

## Price Comparison

Compares prices across providers for equivalent flights.

The system automatically marks the cheapest flight within the same comparison group.

Example:

```json
"is_cheapest": true
```

---

## Best Value Ranking

Flights are ranked using a normalized weighted scoring system based on:

* Price
* Total duration
* Number of stops

This simulates a more realistic flight ranking strategy where cheaper, faster, and more convenient flights receive a better score.

---

## Concurrency

Provider searches are executed concurrently using:

* Goroutines
* Channels

This significantly improves aggregation performance by allowing all providers to be queried in parallel instead of sequentially.

---

## Retry Logic

AirAsia provider simulates random failures.

The system automatically retries failed requests with retry delay handling.

---

## Timeout Handling

Provider aggregation includes timeout handling to prevent slow providers from blocking the entire search process.

---

## Partial Failure Support

If one provider fails or times out, the system still returns results from other providers.

---

## Validation

The system validates normalized flight data before aggregation.

Validation includes:

* Invalid timestamps
* Invalid duration
* Invalid airport codes
* Invalid prices
* Arrival before departure

Invalid flights are automatically skipped.

---

## Response Metadata

Search response includes metadata such as:

* Total results
* Providers queried
* Providers succeeded
* Providers failed
* Search time
* Cache hit status

---

# Architecture Overview

```txt
Providers
   ↓
Normalization
   ↓
Validation
   ↓
Aggregation
   ↓
Search
   ↓
Filtering
   ↓
Ranking
   ↓
Sorting
   ↓
Unified Response
```

---

# Concurrency Design

All provider searches are executed concurrently using goroutines.

Results are collected through channels and aggregated into a unified response.

```txt
Garuda ─┐
Lion ───┼── Concurrent Provider Queries
Batik ──┤
AirAsia ┘
        ↓
Channel Aggregation
        ↓
Merged Flight Results
```

Benefits:

* Faster response time
* Independent provider execution
* Better scalability
* Realistic aggregation behavior

---

# Retry & Timeout Strategy

## Retry Logic

AirAsia provider simulates intermittent provider failures.

If a provider request fails:

* The request is retried up to 3 times
* Retry delay is applied between attempts

---

## Timeout Handling

Aggregation includes timeout handling to avoid waiting indefinitely for slow providers.

Timeout providers are treated as failed providers while successful providers still return results normally.

---

# Flight Normalization

Each provider returns a completely different schema.

The system normalizes all providers into a unified internal structure:

```json
{
  "id": "QZ7250_AirAsia",
  "provider": "AirAsia",
  "flight_number": "QZ7250",
  "departure": {},
  "arrival": {},
  "duration": {},
  "price": {}
}
```

---

# Example Response

```json
{
  "search_criteria": {
    "origin": "CGK",
    "destination": "DPS",
    "departure_date": "2025-12-15"
  },
  "metadata": {
    "total_results": 13,
    "providers_queried": 4,
    "providers_succeeded": 4,
    "providers_failed": 0,
    "search_time_ms": 285,
    "cache_hit": false
  },
  "flights": []
}
```

---

# Project Structure

```txt
flight-project/
│
├── data/
│   ├── airasia_search_response.json
│   ├── batik_air_search_response.json
│   ├── garuda_indonesia_search_response.json
│   ├── lion_air_search_response.json
│   └── expected_result.json
│
├── models/
│   ├── aggregation_result.go
│   ├── filter.go
│   ├── flight.go
│   ├── search.go
│   ├── search_response.go
│   └── sort.go
│
├── providers/
│   ├── airasia.go
│   ├── batik.go
│   ├── garuda.go
│   └── lion.go
│
├── services/
│   ├── aggregator.go
│   ├── comparison.go
│   ├── filter.go
│   ├── ranking.go
│   ├── search.go
│   └── sort.go
│
├── utils/
│   ├── cabin.go
│   ├── duration.go
│   ├── json.go
│   ├── retry.go
│   ├── time.go
│   └── validation.go
│
├── main.go
├── go.mod
└── README.md
```

---

# Design Decisions

## Why Goroutines + Channels?

Go concurrency primitives are lightweight and well-suited for aggregation systems where multiple providers can be queried independently.

---

## Why Local JSON Files?

The assignment explicitly requested mocked provider APIs without actual HTTP requests.

Local JSON files are used to simulate provider responses.

---

## Why Partial Failure Support?

Real-world aggregation systems should continue serving results even when one provider becomes unavailable.

---

## Why Normalized Scoring?

Best-value ranking uses normalized scoring to avoid one metric dominating the ranking system.

This creates a more balanced comparison between:

* Price
* Duration
* Convenience

---

## Why Nested Flight Model?

The nested response structure is more extensible, maintainable, and closer to real-world production APIs.

---

# Assumptions

* Flights with the same route and departure timestamp are considered comparable
* Provider JSON files simulate external airline APIs
* Provider responses may contain inconsistent datetime formats
* Missing timezone values are parsed using Go default parsing behavior

---

# Setup & Run Instructions

## Requirements

* Go 1.22+

---

## 1. Clone Repository

```bash
git clone <repository-url>
```

---

## 2. Navigate To Project

```bash
cd flight-project
```

---

## 3. Install Dependencies

```bash
go mod tidy
```

---

## 4. Run Application

```bash
go run .
```

---

# Author

David
