package collections

import "go.neonxp.dev/extra/utils"

// MergeScalar стабильное слияние двух слайсов скаляров
func MergeScalar[T utils.Scalar](s1, s2 []T) []T {
	i, j := 0, 0
	len1, len2 := len(s1), len(s2)
	result := make([]T, 0, len1+len2)
	for i < len1 && j < len2 {
		if s1[i] < s2[j] {
			result = append(result, s1[i])
			i++
			continue
		}
		result = append(result, s2[j])
		j++
	}
	result = append(result, s1[i:]...)
	result = append(result, s2[j:]...)
	return result
}

// Merge стабильное слияние двух слайсов
func Merge[T utils.Sortable[T]](s1, s2 []T) []T {
	i, j := 0, 0
	len1, len2 := len(s1), len(s2)
	result := make([]T, 0, len1+len2)
	for i < len1 && j < len2 {
		if s1[i].Less(s2[j]) {
			result = append(result, s1[i])
			i++
			continue
		}
		result = append(result, s2[j])
		j++
	}
	result = append(result, s1[i:]...)
	result = append(result, s2[j:]...)
	return result
}
