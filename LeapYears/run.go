package LeapYears

func IsLeapYear(year int) bool {
	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}
