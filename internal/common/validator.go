package common

func ValidateLength(input string, min int, max int) bool {
	if len(input) < min || len(input) > max {
		return false
	}

	return true
}