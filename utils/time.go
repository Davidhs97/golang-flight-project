package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func ParseTimestamp(datetime string) int64 {

	layouts := []string{
		time.RFC3339,
		"2006-01-02T15:04:05-0700", //formatdate batik
		"2006-01-02T15:04:05",      //formatdate lion
	}

	for _, layout := range layouts {

		parsedTime, err :=
			time.Parse(layout, datetime)

		if err == nil {

			return parsedTime.Unix()
		}
	}

	fmt.Println(
		"failed to parse datetime:",
		datetime,
	)

	return 0
}

func FormatDuration(totalMinutes int) string {

	hours := totalMinutes / 60

	minutes := totalMinutes % 60

	return fmt.Sprintf(
		"%dh %dm",
		hours,
		minutes,
	)
}

func ParseTravelTimeToMinutes(travelTime string) int {
	parts := strings.Split(travelTime, " ")

	hours := 0
	minutes := 0

	for _, part := range parts {

		if strings.Contains(part, "h") {
			h := strings.ReplaceAll(part, "h", "")

			hours, _ = strconv.Atoi(h)
		}

		if strings.Contains(part, "m") {
			m := strings.ReplaceAll(part, "m", "")

			minutes, _ = strconv.Atoi(m)
		}
	}

	return (hours * 60) + minutes
}

func ConvertHoursToMinutes(hours float64) int {
	return int(math.Round(hours * 60))
}
