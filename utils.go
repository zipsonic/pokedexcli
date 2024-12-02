package main

import "strconv"

func isNumberInRange(s string, min, max int) bool {
	// Convert the string to an integer
	num, err := strconv.Atoi(s)
	if err != nil {
		// If parsing fails, the string is not a valid number
		return false
	}
	// Check if the number is within the specified range
	return num >= min && num <= max
}
