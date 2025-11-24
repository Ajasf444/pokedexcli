package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		input    string
		expected []byte
	}{
		{
			input:    "www.website.com",
			expected: []byte("data_1"),
		},
		{
			input:    "www.website2.com",
			expected: []byte("data_2"),
		},
	}
	for in, out := range cases {
		t.Run(fmt.Sprintf("Test Case %v", in), func(t *testing.T) {
			// TODO: add tests
		})
	}
}
