package utils

import "strings"

func NormalizeCabinClass(value string) string {

	value = strings.ToUpper(value)

	switch value {

	case "Y", "B", "H", "K", "L", "M", "N", "Q", "S", "T", "V":
		return "economy"

	case "J", "C":
		return "business"

	case "F":
		return "first"

	default:
		return "unknown"
	}
}
