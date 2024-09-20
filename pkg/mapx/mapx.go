package mapx

func ValueOr[K comparable, V any](in map[K]V, key K, fallback V) V {
	if v, ok := in[key]; ok {
		return v
	}
	return fallback
}

func HasKey[K comparable, V any](in map[K]V, key K) bool {
	_, ok := in[key]
	return ok
}
