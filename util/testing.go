package util

import "testing"

func EndTestIfError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Unexpected error, weren't testing for this: %v", err)
	}
}

func AssertEqu(expected interface{}, got interface{}, t *testing.T) {
	if expected != got {
		t.Errorf("Expected '%v', got '%v'", expected, got)
	}
}
