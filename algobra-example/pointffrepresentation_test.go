package main

import (
	"testing"

	"github.com/glycerine/algobra/finitefield"
	"github.com/glycerine/algobra/finitefield/ff"
	"github.com/grbll/go-projective-geometry/pg2"
)

func TestPoint(t *testing.T) {
	field, err := finitefield.Define(7)
	if err != nil {
		t.Fatal(err)
	}

	p := pg2.PointFFRepresentation[ff.Element]{
		X: field.ElementFromUnsigned(1),
		Y: field.ElementFromUnsigned(2),
		Z: field.ElementFromUnsigned(3),
	}

	_ = p
}
