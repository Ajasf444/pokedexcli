package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "www.website.com",
			value: []byte("data_1"),
		},
		{
			key:   "www.website2.com",
			value: []byte("data_2"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.value)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.value) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	const website = "www.example.com"
	cache := NewCache(baseTime)
	cache.Add(website, []byte("data"))
	_, ok := cache.Get(website)
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	time.Sleep(waitTime)
	_, ok = cache.Get(website)
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
