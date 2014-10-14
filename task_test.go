package emru

import "testing"

func TestToggle(t *testing.T) {
	tests := []struct {
		s      status
		expect status
	}{
		{true, false},
		{false, true},
	}
	for i, test := range tests {
		test.s.toggle()
		if test.s != test.expect {
			t.Errorf("Test %d: Expected %s after toggle but got %s", i, test.expect, test.s)
		}
	}
}
