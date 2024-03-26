package assert

import "testing"

func Equals[T comparable](t *testing.T, left, right T) {
	if left != right {
		t.Errorf("Left %#v, Right %#v", left, right)
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error %v", err)
	}
}
