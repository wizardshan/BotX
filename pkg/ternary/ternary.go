package ternary

func Do[T any](condition bool, ifOutput T, elseOutput T) T {
	if condition {
		return ifOutput
	}

	return elseOutput
}
