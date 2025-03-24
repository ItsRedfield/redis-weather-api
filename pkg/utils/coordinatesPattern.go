package utils

import (
	"regexp"
)

func CoordinatePattern(coordinate string) (bool, error) {
	validPattern := `^[-]?\d+(\.\d+)?$`
	return regexp.MatchString(validPattern, coordinate)
}
