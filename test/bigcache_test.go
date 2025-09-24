package test

import (
	"testing"

	"github.com/allegro/bigcache/v3"
)

func TestBigCache(t *testing.T) {
	cache, _ := bigcache.New(t.Context(), bigcache.Config{})

	cache.Set("my-unique-key", []byte("value"))
	entry, _ := cache.Get("my-unique-key")

	if string(entry) != "value" {
		t.Errorf("expected value to be 'value', got '%s'", entry)
	}

	t.Log("BigCache test passed")
}
