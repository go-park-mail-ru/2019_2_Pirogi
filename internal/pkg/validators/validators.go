package validators

import "regexp"

var emailPattern = regexp.MustCompile(`^([a-z0-9_-]+\.)*[a-z0-9_-]+@[a-z0-9_-]+(\.[a-z0-9_-]+)*\.[a-z]{2,6}$`)

func ValidateEmail(email string) bool {
	if emailPattern.MatchString(email) {
		return true
	}
	return false
}
