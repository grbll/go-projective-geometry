package ffrepresentation_test

import (
	"testing"

	ffr "github.com/grbll/go-projective-geometry/pg2/ffrepresentation"
)

func TestCoordinatesFrom(t *testing.T) {
	want := []ffr.Coordinates{
		{X: 0, Y: 0, Z: 1},
		{X: 0, Y: 1, Z: 0},
		{X: 0, Y: 1, Z: 1},
		{X: 1, Y: 0, Z: 0},
		{X: 1, Y: 0, Z: 1},
		{X: 1, Y: 1, Z: 0},
		{X: 1, Y: 1, Z: 1},
	}

	var got []ffr.Coordinates
	for c := range ffr.CoordinatesFrom(ffr.Coordinates{X: 0, Y: 0, Z: 1}, 2) {
		got = append(got, c)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("coordinate %d = %v, want %v", i, got[i], want[i])
		}
	}
}
