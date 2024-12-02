package main

import "strconv"

func isNumberInRange(s string, min, max int) (int, bool) {
	// Convert the string to an integer
	num, err := strconv.Atoi(s)
	if err != nil {
		// If parsing fails, the string is not a valid number
		return 0, false
	}
	// Check if the number is within the specified range
	return num, (num >= min && num <= max)
}
