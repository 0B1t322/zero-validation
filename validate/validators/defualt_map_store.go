package validators

type defaultMapStore struct {
	store map[string]any
}

func (d *defaultMapStore) Get(key string) any {
	return d.store[key]
}

func (d *defaultMapStore) Set(key string, v any) {
	d.store[key] = v
}

func (d *defaultMapStore) All() []any {
	all := make([]any, 0, len(d.store))
	for _, v := range d.store {
		all = append(all, v)
	}
	return all
}

// NewDefaultMapStore return default map store
// it's not for concurrent write
func NewDefaultMapStore() ValidatorStore {
	return &defaultMapStore{
		store: make(map[string]any),
	}
}
