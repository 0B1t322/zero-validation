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

func (d *defaultMapStore) Range(f func(key string, value any) bool) {
	for k, v := range d.store {
		if !f(k, v) {
			break
		}
	}
}

func DefaultMapStore() ValidatorStore {
	return &defaultMapStore{
		store: make(map[string]any),
	}
}
