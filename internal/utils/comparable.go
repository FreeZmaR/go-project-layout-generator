package utils

func OneOf[T comparable](value T, values ...T) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}

	return false
}
