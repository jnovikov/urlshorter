package keystorage

type StorageWrapper struct {
	KeyStorage
}

func (st *StorageWrapper) Exists(key string) bool {
	result := st.Get(key)
	return result != nil
}

func (st *StorageWrapper) GetString(key string) string {
	result := st.Get(key)
	if result != nil {
		switch v := result.(type) {
		case string:
			return v
		}
	}
	return ""
}

func NewWrapper(storage KeyStorage) *StorageWrapper {
	return &StorageWrapper{storage}
}
