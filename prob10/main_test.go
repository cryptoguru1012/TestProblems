package main

import (
	"math"
	"math/big"
	"testing"
)

// for a valid return value.
func TestPrimeFilterForZero(t *testing.T) {
	want := big.NewInt(0)
	_, sum, err := PrimeFilter(0)
	if want.Cmp(sum) != 0 || err {
		t.Fatalf(`PrimeFilter(0) = %d, want match for %d`, sum, want)
	}
}

// for a vaild return value
func TestPrimeFilterFor2Milion(t *testing.T) {
	want := big.NewInt(142913828922)
	_, sum, err := PrimeFilter(2000000)
	if want.Cmp(sum) != 0 || err {
		t.Fatalf(`PrimeFilter(2000000) = %d, want match for %d`, sum, want)
	}
}

// for invalid input value
func TestPrimeFilterForMaxInt64(t *testing.T) {
	want := big.NewInt(math.MaxInt64)
	_, sum, err := PrimeFilter(math.MaxInt64)
	if err {
		t.Fatalf(`(bad input case)PrimeFilter(MaxInt64:%d) = %d, want match for %d`, math.MaxInt64, sum, want)
	}
}
