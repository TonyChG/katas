package FizzBuzz

import "testing"

func AssertFizzBuzz(t *testing.T, input int, expectedOutput string) {
	output := Run(input)
	if output != expectedOutput {
		t.Errorf("FizzBuzz.Run(%d) = %s; want %s", input, output, expectedOutput)
	}
}

func TestEmptyOutput(t *testing.T) {
	AssertFizzBuzz(t, 1, "")
}

func TestThreeToFizz(t *testing.T) {
	AssertFizzBuzz(t, 3, "Fizz")
}

func TestFiveToFizz(t *testing.T) {
	AssertFizzBuzz(t, 5, "Buzz")
}

func TestFifteenToFizzBuzz(t *testing.T) {
	AssertFizzBuzz(t, 15, "FizzBuzz")
}

func TestThreeDigitToFizz(t *testing.T) {
	AssertFizzBuzz(t, 43, "Fizz")
}

func TestFiveDigitToBuzz(t *testing.T) {
	AssertFizzBuzz(t, 52, "Buzz")
}

func TestFiveThreeDigitToFizzBuzz(t *testing.T) {
	AssertFizzBuzz(t, 53, "FizzBuzz")
}
