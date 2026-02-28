package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key string
		value []byte
	} {
		{
		key: "https://testwebsiteone.com",
		value: []byte("this is one case"),
		},
		{
		key: "https://anotherwebsite.com",
		value: []byte("this is the first value in another case"),
		},
	}

	const interval = 5 * time.Second
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.value)
			
			value, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key in cache")
				return
			}

			if string(value) != string(c.value) {
				t.Errorf("Unexpected value found in cache")
				return
			}
		})
	}
}