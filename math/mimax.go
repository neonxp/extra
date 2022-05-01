package math

import "go.neonxp.dev/extra/utils"

// MinScalar возвращает минимальное из переданных скалярных значений
func MinScalar[T utils.Scalar](a ...T) T {
	i := -1
	for j, v := range a {
		if i == -1 {
			i = j
			continue
		}
		if v < a[i] {
			i = j
		}
	}
	return a[i]
}

// MaxScalar возвращает максимальное из переданных скалярных значений
func MaxScalar[T utils.Scalar](a ...T) T {
	i := -1
	for j, v := range a {
		if i == -1 {
			i = j
			continue
		}
		if v > a[i] {
			i = j
		}
	}
	return a[i]
}

// Min возвращает минимальное из переданных Sortable значений
func Min[T utils.Sortable[T]](a ...T) T {
	i := -1
	for j, v := range a {
		if i == -1 {
			i = j
			continue
		}
		if v.Less(a[i]) {
			i = j
		}
	}
	return a[i]
}

// Max возвращает максимальное из переданных Sortable значений
func Max[T utils.Sortable[T]](a ...T) T {
	i := -1
	for j, v := range a {
		if i == -1 {
			i = j
			continue
		}
		if a[i].Less(v) {
			i = j
		}
	}
	return a[i]
}
