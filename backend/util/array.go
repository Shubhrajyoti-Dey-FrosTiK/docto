package util

import "fmt"

func ArrayContains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func UniqueArray[I any](arrayI *[]I) *[]I {
	var hashMap = make(map[string]bool)
	var result = make([]I, 0)

	for _, element := range *arrayI {
		key := fmt.Sprintf("%v", element)
		if hashMap[key] == false {
			result = append(result, element)
			hashMap[key] = true
		}
	}

	return &result
}

func ArrayPaginate[T any](arrayT *[]T, pageSize int, pageNumber int) *[]T {
	result := make([]T, 0)
	startIdx := pageSize * (pageNumber - 1)

	for itr := startIdx; itr < len(*arrayT) && itr < startIdx+pageSize; itr++ {
		result = append(result, (*arrayT)[itr])
	}

	return &result
}
