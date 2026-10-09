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
		in   ffr.Coordinates
		want ffr.Coordinates
	}{
		{
			name: "x nonzero",
			in:   ffr.Coordinates{X: 2, Y: 3, Z: 4},
			want: ffr.Coordinates{X: 1, Y: 4, Z: 2},
		},
		{
			name: "x zero y nonzero",
			in:   ffr.Coordinates{X: 0, Y: 2, Z: 3},
			want: ffr.Coordinates{X: 0, Y: 1, Z: 4},
		},
		{
			name: "x and y zero",
			in:   ffr.Coordinates{X: 0, Y: 0, Z: 3},
			want: ffr.Coordinates{X: 0, Y: 0, Z: 1},
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

func TestNew(t *testing.T) {
	gf5, err := finitefield.Define(5)
	if err != nil {
		t.Fatal(err)
	}

	field := New(gf5)
	plane := ffr.New(field)

	// q^3 - 1 nonzero coordinate triples are used as map keys.
	if got, want := len(plane.Points), 5*5*5-1; got != want {
		t.Errorf("got %d point keys, want %d", got, want)
	}

	if got, want := len(plane.Lines), 5*5*5-1; got != want {
		t.Errorf("got %d line keys, want %d", got, want)
	}

	// There should be q^2 + q + 1 distinct point/line objects.
	points := make(map[*ffr.Point[AlgobraElement]]struct{})
	lines := make(map[*ffr.Line[AlgobraElement]]struct{})

	for c, point := range plane.Points {
		points[point] = struct{}{}

		if plane.Canonical(c) != point.Rep {
			t.Errorf("point at %v has representative %v, but should %v", c, point.Rep, plane.Canonical(c))
		}
	}

	for c, line := range plane.Lines {
		lines[line] = struct{}{}

		if plane.Canonical(c) != line.Rep {
			t.Errorf("line at %v has representative %v", c, line.Rep)
		}
	}

	if got, want := len(points), 5*5+5+1; got != want {
		t.Errorf("got %d distinct points, want %d", got, want)
	}

	if got, want := len(lines), 5*5+5+1; got != want {
		t.Errorf("got %d distinct lines, want %d", got, want)
	}
}

func TestGenerateLineInclusions(t *testing.T) {
	gf5, err := finitefield.Define(5)
	if err != nil {
		t.Fatal(err)
	}

	plane := ffr.New(New(gf5))

	tests := []struct {
		name string
		line ffr.Coordinates
		want []ffr.Coordinates
	}{
		{
			name: "z=0, y=0",
			line: ffr.Coordinates{X: 1, Y: 0, Z: 0},
			want: []ffr.Coordinates{
				{X: 0, Y: 0, Z: 1},
				{X: 0, Y: 1, Z: 0},
				{X: 0, Y: 1, Z: 1},
				{X: 0, Y: 1, Z: 2},
				{X: 0, Y: 1, Z: 3},
				{X: 0, Y: 1, Z: 4},
			},
		},
		{
			name: "z=0, y!=0",
			line: ffr.Coordinates{X: 1, Y: 2, Z: 0},
			want: []ffr.Coordinates{
				{X: 0, Y: 0, Z: 1},
				{X: 1, Y: 2, Z: 0},
				{X: 1, Y: 2, Z: 1},
				{X: 1, Y: 2, Z: 2},
				{X: 1, Y: 2, Z: 3},
				{X: 1, Y: 2, Z: 4},
			},
		},
		{
			name: "z!=0",
			line: ffr.Coordinates{X: 1, Y: 2, Z: 1},
			want: []ffr.Coordinates{
				{X: 0, Y: 1, Z: 3},
				{X: 1, Y: 0, Z: 4},
				{X: 1, Y: 1, Z: 2},
				{X: 1, Y: 2, Z: 0},
				{X: 1, Y: 3, Z: 3},
				{X: 1, Y: 4, Z: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plane.GenerateLineInclusions(tt.line)

			if len(*got) != len(tt.want) {
				t.Fatalf("got %d points, want %d", len(*got), len(tt.want))
			}

			for i, want := range tt.want {
				got := (*got)[i].Rep

				if got != want {
					t.Errorf("point %d = %v, want %v", i, got, want)
				}
			}
		})
	}
}
