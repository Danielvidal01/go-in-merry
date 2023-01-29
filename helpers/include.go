package helpers

func Includes[T comparable](values []T, value T) bool {
	for _, val := range values {
		if val == value {
			return true
		}
	}

	return false
}
