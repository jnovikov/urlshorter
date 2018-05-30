package keystorage


type MockKeyStorage struct {
	m map[string]interface{}
}

func (ms *MockKeyStorage) Get(key string) interface{} {
	return ms.m[key]
}

func (ms *MockKeyStorage) Set(key string, value interface{}) {
	ms.m[key] = value
}


func NewMockStorage() *MockKeyStorage {
	s := new(MockKeyStorage)
	s.m = make(map[string]interface{})
	return s
}