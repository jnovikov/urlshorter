package keystorage

import (
	"testing"
)

func TestStorage_NotExists(t *testing.T) {
	m := NewMockStorage()
	res := m.Get("not_found")
	if res != nil {
		t.Error("Failed to check what key not found")
	}
}

func TestStorageWrapper_Exists(t *testing.T) {
	m := NewMockStorage()
	w := StorageWrapper{m}
	if w.Exists("not_found") {
		t.Error("Failed to check what key not found")
	}
}

func TestStorageWrapper_GetString(t *testing.T) {
	m := NewMockStorage()
	w := StorageWrapper{m}
	w.Set("asd", "asd")

	if w.GetString("asd") != "asd" {
		t.Error("Failed to get string by key")
	}
}

func TestStorageWrapper_GetBadString(t *testing.T) {
	m := NewMockStorage()
	w := StorageWrapper{m}
	w.Set("asd", 1)
	res := w.GetString("asd")
	if res != "" {
		t.Error("Failed to get empty string then type error")
	}

}
