package slices

func Associate[T any, K comparable, V any](collection []T, f func(item T) (K, V)) map[K]V {
	m := make(map[K]V, len(collection))

	for _, item := range collection {
		key, value := f(item)
		m[key] = value
	}

	return m
}
