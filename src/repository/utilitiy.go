package repository

import "url_manager/utility"

func IndexOf[T comparable](src []T, target T) int {
	return utility.IndexOf(src, target)
}

func FindByFunc[T any](src []T, fc func(t T) bool) int {
	return utility.IndexOfFunc(src, fc)
}

func Remove[T any](src []T, i int) []T {
	return utility.Remove(src, i)
}
