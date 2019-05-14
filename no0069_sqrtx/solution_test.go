package no0069_sqrtx

import "testing"

func TestMySqrt(test *testing.T) {

	if mySqrt(4) != 2 {
		test.Error("error", 4, mySqrt(4))
	}

	if mySqrt(8) != 2 {
		test.Error("error", 8, mySqrt(8))
	}
}
