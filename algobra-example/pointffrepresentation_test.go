package main

import (
	"testing"

	"github.com/glycerine/algobra/finitefield"
	ffr "github.com/grbll/go-projective-geometry/pg2/ffrepresentation"
)

func TestCanonical(t *testing.T) {
	gf5, err := finitefield.Define(5)
	if err != nil {
		t.Fatal(err)
	}
	afield := New(gf5)
	plane := ffr.FFProjectivePlane[AlgobraElement]{
		Field: afield,
	}

	tests := []struct {
		name string
		in   ffr.Coordinates[AlgobraElement]
		want ffr.Coordinates[AlgobraElement]
	}{
		{
			name: "x nonzero",
			in:   ffr.Coordinates[AlgobraElement]{X: 2, Y: 3, Z: 4},
			want: ffr.Coordinates[AlgobraElement]{X: 1, Y: 4, Z: 2},
		},
		{
			name: "x zero y nonzero",
			in:   ffr.Coordinates[AlgobraElement]{X: 0, Y: 2, Z: 3},
			want: ffr.Coordinates[AlgobraElement]{X: 0, Y: 1, Z: 4},
		},
		{
			name: "x and y zero",
			in:   ffr.Coordinates[AlgobraElement]{X: 0, Y: 0, Z: 3},
			want: ffr.Coordinates[AlgobraElement]{X: 0, Y: 0, Z: 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plane.Canonical(tt.in)
			if got != tt.want {
				t.Errorf("Canonical(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
