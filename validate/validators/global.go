package validators

var globalMapStore = NewConcurrentMapStore()

func GlobalMapStore() ValidatorStore {
	return globalMapStore
}
