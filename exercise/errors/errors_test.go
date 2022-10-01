package timeparse

import (
	"fmt"
	"testing"
)

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"1:3:44", true},
		{"bad", false},
		{"1:-3:44", false},
		{"0:59:59", true},
		{"", false},
		{"11:12", false},
		{"aa:bb:cc", false},
		{"5:23", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		fmt.Println(err)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}
