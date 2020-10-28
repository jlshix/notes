package iteration

import "testing"

func TestRepeat(t *testing.T) {
	rv := Repeat("x", 5)
	expected := "xxxxx"
	if rv != expected {
		t.Errorf("expected %s but got %s", expected, rv)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("n", 5)
	}
}
