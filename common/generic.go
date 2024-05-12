package common

func IfError[T any](Value T, err error, Default T) T {
	if err != nil {
		return Default
	}
	return Value
}
