package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := strings.Repeat("a", 4)
	expected := "aaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 4)
	}
}

func ExampleRepeat() {
	repeated := strings.Repeat("b", 6)
	fmt.Println(repeated)
	// Output: bbbbbb
}
