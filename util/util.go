package util

func IsNumber(r rune) bool {

	if r >= 48 && r <= 57 {
		return true
	}
	return false
}
