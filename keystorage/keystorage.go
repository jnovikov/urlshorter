package keystorage

type KeyStorage interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}