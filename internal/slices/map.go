package slices

func Map[T any, V any](collection []T, f func(int, T) V) []V {
	out := make([]V, 0, len(collection))
	for i, item := range collection {
		out = append(out, f(i, item))
	}

	return out
}
