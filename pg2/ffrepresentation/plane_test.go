package ffrepresentation_test

import (
	"reflect"
	"testing"

	ffr "github.com/grbll/go-projective-geometry/pg2/ffrepresentation"
)

func TestCoordinatesFrom(t *testing.T) {
	k := uint(3)

	expected := []ffr.Coordinates{
		{X: 0, Y: 0, Z: 1},
		{X: 0, Y: 0, Z: 2},
		{X: 0, Y: 1, Z: 0},
		{X: 0, Y: 1, Z: 1},
		{X: 0, Y: 1, Z: 2},
		{X: 0, Y: 2, Z: 0},
		{X: 0, Y: 2, Z: 1},
		{X: 0, Y: 2, Z: 2},
		{X: 1, Y: 0, Z: 0},
		{X: 1, Y: 0, Z: 1},
		{X: 1, Y: 0, Z: 2},
		{X: 1, Y: 1, Z: 0},
		{X: 1, Y: 1, Z: 1},
		{X: 1, Y: 1, Z: 2},
		{X: 1, Y: 2, Z: 0},
		{X: 1, Y: 2, Z: 1},
		{X: 1, Y: 2, Z: 2},
		{X: 2, Y: 0, Z: 0},
		{X: 2, Y: 0, Z: 1},
		{X: 2, Y: 0, Z: 2},
		{X: 2, Y: 1, Z: 0},
		{X: 2, Y: 1, Z: 1},
		{X: 2, Y: 1, Z: 2},
		{X: 2, Y: 2, Z: 0},
		{X: 2, Y: 2, Z: 1},
		{X: 2, Y: 2, Z: 2},
	}

	var got []ffr.Coordinates

	for c := range ffr.CoordinatesFrom(
		ffr.Coordinates{X: 0, Y: 0, Z: 1},
		k,
	) {
		got = append(got, c)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, want %v", got, expected)
	}
}
