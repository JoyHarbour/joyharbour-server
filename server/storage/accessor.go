package storage

var storage Storage

func GetStorage() Storage {
	if storage == nil {
		panic("GetStorage() called before LoadStorage()")
	}

	return storage
}
