package FizzBuzz

import "testing"

func TestEmptyOutput(t *testing.T) {
	input := 1
	expectedOutput := ""
	output := Run(input)
	if output != expectedOutput {
		t.Errorf("FizzBuzz.Run(%d) = %s; want %s", input, output, expectedOutput)
	}
}

func TestThreeToFizz(t *testing.T) {
	input := 3
	expectedOutput := "Fizz"
	output := Run(input)
	if output != expectedOutput {
		t.Errorf("FizzBuzz.Run(%d) = %s; want %s", input, output, expectedOutput)
	}
}

func TestFiveToFizz(t *testing.T) {
	input := 5
	expectedOutput := "Buzz"
	output := Run(input)
	if output != expectedOutput {
		t.Errorf("FizzBuzz.Run(%d) = %s; want %s", input, output, expectedOutput)
	}
}

func TestFifteenToFizzBuzz(t *testing.T) {
	input := 15
	expectedOutput := "FizzBuzz"
	output := Run(input)
	if output != expectedOutput {
		t.Errorf("FizzBuzz.Run(%d) = %s; want %s", input, output, expectedOutput)
	}
}
