package utils

// Ptr получение указателя от произвольного значения, например, ToPtr(true) -> *bool(true)
func Ptr[T any](t T) *T {
	return &t
}
