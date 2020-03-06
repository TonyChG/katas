package FizzBuzz

import (
	"strconv"
	"strings"
)

func Run(number int) (output string) {
	output = ""
	converted := strconv.Itoa(number)

	if (number%3 == 0 && number%5 == 0) || (strings.ContainsAny(converted, "3") && strings.ContainsAny(converted, "5")) {
		output = "FizzBuzz"
	} else if number%3 == 0 || strings.ContainsAny(converted, "3") {
		output = "Fizz"
	} else if number%5 == 0 || strings.ContainsAny(converted, "5") {
		output = "Buzz"
	}
	return
}
