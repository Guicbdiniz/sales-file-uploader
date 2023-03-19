package utils

import "testing"

// CheckTestError is a test helper to confirm that an error is nil.
//
// If error is not nil, test will fail.
func CheckTestError(t *testing.T, err error, message string) {
	t.Helper()

	if err != nil {
		t.Fatalf("%s: %s", message, err.Error())
	}
}

// AssertEqual compares two values, failing if they are not equal.
func AssertEqual[V comparable](t *testing.T, expected V, actual V, message string) {
	t.Helper()

	if expected != actual {
		t.Fatalf("%s. Expected %v, got %v", message, expected, actual)
	}
}
