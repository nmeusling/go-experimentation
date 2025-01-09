package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	expectedKey := "http://www.example.com"
	expectedValue := []byte("test")
	cache := NewCache(5 * time.Second)
	cache.Add(expectedKey, expectedValue)
	val, ok := cache.Get(expectedKey)
	if !ok {
		t.Errorf("Expected to find key %v among", expectedKey)
		return
	}
	if string(val) != string(expectedValue) {
		t.Errorf("Expected to find value %v but found %v", string(expectedValue), string(val))
		return
	}
}

func TestGetNonexistent(t *testing.T) {
	nonExistentKey := "test"
	cache := NewCache(5 * time.Second)
	val, ok := cache.Get(nonExistentKey)
	if ok {
		t.Errorf("Did not expect to find value %v for key %v", val, nonExistentKey)
		return
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const key = "http://www.example.com"
	cache := NewCache(baseTime)
	cache.Add(key, []byte("test"))

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key %v", key)
		return
	}
	time.Sleep(baseTime * 2)

	_, ok = cache.Get(key)
	if ok {
		t.Errorf("expected not to find key %v", key)
		return
	}

}
