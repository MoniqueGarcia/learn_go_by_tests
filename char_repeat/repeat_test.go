package iteration

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if result != expected {
		t.Errorf("result is %q, expected %q", result, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}