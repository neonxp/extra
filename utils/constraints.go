package utils

// Scalar соответствует скалярам над которыми определены операции сравнения
type Scalar interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// Sortable соответствует типам над которыми определена операция Less
type Sortable[T any] interface {
	Less(T) bool
}
