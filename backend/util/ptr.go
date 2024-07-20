package util

func ToPtr[T any](t T) *T {
	return &t
}

func NilPtr[T any]() *T {
	return nil
}
