package LeapYears

import "strconv"

func IsLeapYear(year int) bool {
	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

func IsLeapYearString(year string) (bool, error) {
	c, err := strconv.Atoi(year)
	return IsLeapYear(c), err
}
