package utility

func IndexOf[T comparable](src []T, target T) int {
	for i, item := range src {
		if item == target {
			return i
		}
	}
	return -1
}

func IndexOfFunc[T any](src []T, fn func(t T) bool) int {
	for i, item := range src {
		if fn(item) {
			return i
		}
	}
	return -1
}

func Remove[T any](src []T, i int) []T {
	return append(src[:i], src[i+1:]...)
}
