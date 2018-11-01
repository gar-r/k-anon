package util

import "testing"

func Test_MinMax(t *testing.T) {
	if !(MinInt < MinUint && MinUint < MaxInt) {
		t.Errorf("Incorrect const definiton")
	}
}
