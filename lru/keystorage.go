package lru

import "github.com/hashicorp/golang-lru"

type KeyStorage struct {
	l *lru.Cache
}

func (st *KeyStorage) Get(key string) interface{} {
	if res, ok := st.l.Get(key); ok {
		return res
	}
	return nil
}

func (st *KeyStorage) Set(key string, value interface{}) {
	st.l.Add(key, value)
}

func Storage(size int) *KeyStorage {
	l, _ := lru.New(size)
	st := new(KeyStorage)
	st.l = l
	return st
}
