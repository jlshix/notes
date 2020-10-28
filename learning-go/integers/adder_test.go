package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	rv := Add(1, 2)
	expected := 3
	if rv != expected {
		t.Errorf("Add(1, 2) expected %d, got %d", expected, rv)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
