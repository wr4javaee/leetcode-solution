package no0374_guess_number_higher_or_lower

import "testing"

func TestGuessNumber(test *testing.T) {

	if guessNumber(10, 6) != 6 {
		test.Error("error", 10, 6, guessNumber(10, 6))
	}
}
