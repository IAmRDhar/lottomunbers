package main

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

func within(a, b, maxDiff float64) bool {
	absDiff := math.Abs(a - b)
	return absDiff < maxDiff
}

func TestLottoNumbers(t *testing.T) {
	list := lottoNumbers(numberToGenerate)

	const expectedNumberLength = 7

	// list should generated the correct number of numbers
	if len(list) != numberToGenerate {
		t.Errorf("len(list) is %v, not %v", len(list), numberToGenerate)
	}

	// lotto numbers are 7 long
	var withIncorrectLength []int
	for i, number := range list {
		if len(number) != expectedNumberLength {
			withIncorrectLength = append(withIncorrectLength, i)
		}
	}
	if len(withIncorrectLength) != 0 {
		t.Errorf("%v lotto numbers including index %v do not have the expected length of %d", len(withIncorrectLength), withIncorrectLength[0], expectedNumberLength)
		t.FailNow() // later tests rely on the numbers existing
	}

	// convert to form needed by stat
	// also grab max, min
	x := make([]float64, numberToGenerate*7)
	max := list[0][0]
	min := max
	for i := range x {
		j := i / 7
		k := i % 7

		n := list[j][k]

		x[i] = float64(n)

		if n > max {
			max = n
		}

		if n < min {
			min = n
		}
	}

	// check max and min
	if max > 48 {
		t.Errorf("max is %v", max)
	}

	if min < 0 {
		t.Errorf("min is %v", min)
	}

	// check distribution properties
	maxDiff := 0.02
	uniform := distuv.Uniform{
		Max: 49,
	}

	stdDev := stat.StdDev(x, nil)
	if !within(stdDev, uniform.StdDev(), maxDiff) {
		t.Errorf("std dev not within %v of uniform (%v): %v", maxDiff, uniform.StdDev(), stdDev)
	}

	skew := stat.Skew(x, nil)
	if !within(skew, uniform.Skewness(), maxDiff) {
		t.Errorf("skewness not within %v of uniform (%v): %v", maxDiff, uniform.Skewness(), skew)
	}

	exk := stat.ExKurtosis(x, nil)
	if !within(exk, uniform.ExKurtosis(), maxDiff) {
		t.Errorf("ExKurtosis not within %v of uniform (%v): %v", maxDiff, uniform.ExKurtosis(), exk)
	}
}
