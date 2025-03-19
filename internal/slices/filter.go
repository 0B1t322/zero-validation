package slices

func Filter[V any](collection []V, predicate func(item V) bool) []V {
	result := make([]V, 0, len(collection))

	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
