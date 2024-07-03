package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	if cache.cache == nil {
		t.Error("Cache is nil")
	}
}
func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 50)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{

			inputKey: "key1",
			inputVal: []byte("val1"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		actual, ok := cache.Get(c.inputKey)
		if !ok {
			t.Errorf("%s not found", c.inputKey)
		}
		if string(actual) != string(c.inputVal) {
			t.Errorf("Values don't match: %s vs %s", string(actual), string(c.inputVal))
		}
	}

}

func TestReap(t *testing.T) {
    interval := time.Millisecond * 10
    cache := NewCache(interval)
    keyOne := "key1"
    cache.Add(keyOne,[]byte("val1"))

    time.Sleep(interval + time.Millisecond*2)

    _, ok := cache.Get(keyOne)
    if ok {
        t.Errorf("%s should have been reaped",keyOne)
    }
}
func TestReapFail(t *testing.T) {
    interval := time.Millisecond * 10
    cache := NewCache(interval)
    keyOne := "key1"
    cache.Add(keyOne,[]byte("val1"))

    time.Sleep(interval / 2)

    _, ok := cache.Get(keyOne)
    if !ok {
        t.Errorf("%s should not have been reaped",keyOne)
    }
}
