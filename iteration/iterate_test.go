package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	out := Repeat("a", 5)
	expected := "aaaaa"

	if out != expected {
		t.Errorf("expected %q but got %q", expected, out)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	out := Repeat("a", 10)
	fmt.Println(out)
	// Output: aaaaaaaaaa
}
