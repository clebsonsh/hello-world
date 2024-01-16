package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})
	t.Run("repeat 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})
	t.Run("repeat 3 times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
		}
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
