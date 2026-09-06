package main

import (
	"testing"

	"github.com/glycerine/algobra/finitefield"
	"github.com/glycerine/algobra/finitefield/ff"
	ffr "github.com/grbll/go-projective-geometry/pg2/ffrepresentation"
)

func TestPoint(t *testing.T) {
	field, err := finitefield.Define(7)
	if err != nil {
		t.Fatal(err)
	}

	p := ffr.Point[ff.Element]{
		X: field.ElementFromUnsigned(1),
		Y: field.ElementFromUnsigned(2),
		Z: field.ElementFromUnsigned(3),
	}

	_ = p
}
