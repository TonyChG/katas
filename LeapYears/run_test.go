package LeapYears

import "testing"

func AssertLeapYears(t *testing.T, input int, expectedOutput bool) {
	output := IsLeapYear(input)
	if output != expectedOutput {
		t.Errorf("LeapYears.IsLeapYear(%d) = %v; want %v", input, output, expectedOutput)
	}
}

func TestDivisibleByFourHundred(t *testing.T) {
	AssertLeapYears(t, 2000, true)
}

func TestDivisibleByOneHundredNotFourHundred(t *testing.T) {
	AssertLeapYears(t, 1700, false)
	AssertLeapYears(t, 1800, false)
	AssertLeapYears(t, 1900, false)
	AssertLeapYears(t, 2100, false)
}

func TestDivisibleByFourNotOneHundred(t *testing.T) {
	AssertLeapYears(t, 2008, true)
	AssertLeapYears(t, 2012, true)
	AssertLeapYears(t, 2016, true)
}

func TestDivisibleNotByFour(t *testing.T) {
	AssertLeapYears(t, 2017, false)
	AssertLeapYears(t, 2018, false)
	AssertLeapYears(t, 2019, false)
}
