package FizzBuzz

func Run(number int) (output string) {
	output = ""
	if number%3 == 0 && number%5 == 0 {
		output = "FizzBuzz"
	} else if number%3 == 0 {
		output = "Fizz"
	} else if number%5 == 0 {
		output = "Buzz"
	}
	return
}
