package helpers

// Must panics with the provided error if it is not nil.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// MustValue returns the provided value if the provided error is nil,
// otherwise it panics with the provided error.
func MustValue[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
