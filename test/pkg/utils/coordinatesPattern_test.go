package utils_test

import (
	"cloudflare-challenge-weaher-api/pkg/utils"
	"testing"
)

func TestCoordinatePattern(t *testing.T) {
	tests := []struct {
		coordinate string
		expected   bool
	}{
		{"45.0", true},
		{"-45.0", true},
		{"90", true},
		{"-90", true},
		{"180.123456", true},
		{"-180.123456", true},
		{"abc", false},
		{"45.0.0", false},
		{"-45.0.0", false},
		{"", false},
	}

	for _, test := range tests {
		result, err := utils.CoordinatePattern(test.coordinate)
		if err != nil {
			t.Errorf("Unexpected error for coordinate %s: %v", test.coordinate, err)
		}
		if result != test.expected {
			t.Errorf("For coordinate %s, expected %v but got %v", test.coordinate, test.expected, result)
		}
	}
}
