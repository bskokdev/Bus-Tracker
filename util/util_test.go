package util

import (
	"testing"
)

// TestGetDistanceKm tests the GetDistanceInKm function
func TestGetDistanceKm(t *testing.T) {
	// Test data
	lat1 := 60.205183
	lon1 := 25.145534

	lat2 := 60.203086
	lon2 := 25.141794

	// Expected result
	expected := 0.3115538933895413

	// Get distance
	distance := GetDistanceInKm(lat1, lon1, lat2, lon2)

	// Check if result is as expected
	if distance != expected {
		t.Errorf("Expected %v, got %v", expected, distance)
	}
}

// TestGetPageOffset tests the GetPageOffset function
func TestGetPageOffset(t *testing.T) {
	// Test data
	page := 2
	pageSize := 20

	// Expected result
	expected := 20

	// Get offset
	offset := GetPageOffset(page, pageSize)

	// Check if result is as expected
	if offset != expected {
		t.Errorf("Expected %v, got %v", expected, offset)
	}
}
