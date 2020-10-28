package hello

import "testing"

// TestHello test for hello
func TestHello(t *testing.T) {
	data := []struct {
		language string
		name     string
		result   string
	}{
		{language: "english", name: "leo", result: "Hello, leo"},
		{language: "chinese", name: "leo", result: "你好, leo"},
	}

	for _, item := range data {
		got := Hello(item.language, item.name)
		if got != item.result {
			t.Errorf("Hello(%s, %s) expected %s, got %s", item.language, item.name, item.result, got)
		}
	}
}
