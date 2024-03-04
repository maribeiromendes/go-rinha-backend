package testutil

import "testing"

func AssertEquals[T int | string](actual T, expected T, t *testing.T) {
	if actual != expected {
		t.Errorf("assert equal failed: actual: %v expected: %v", actual, expected)
	}
}
