package util

import "testing"

func EndTestIfError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Unexpected error, weren't testing for this: %v", err)
	}
}

func AssertEqu(a interface{}, b interface{}, t *testing.T) {
	if a != b {
		t.Errorf("Expected '%v', got '%v'", a, b)
	}
}
