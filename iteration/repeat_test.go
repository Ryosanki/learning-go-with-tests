package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	t.Run("Repeat(\"a\", 5)", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat(\"x\", 10)", func(t *testing.T) {
		repeated := Repeat("x", 10)
		expected := "xxxxxxxxxx"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Repeat(\"a\", 0)", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
