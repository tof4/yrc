package validator

func ValidateLength(input string, min int, max int) bool {
	if len(input) < min || len(input) > max {
		return true
	}

	return false
}
