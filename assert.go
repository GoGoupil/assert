package assert

import "testing"

func AssertEqual(t *testing.T, val interface{}, exp interface{}) {
	if val != exp {
		t.Errorf("Expected %v, got %v.", exp, val)
	}
}
