package iteration

import (
	"fmt"
	"testing"
)

func TestCharacterCount(t *testing.T) {
	count := CharacterCount("my string")
	expected := 9

	if count != expected {
		t.Errorf("Expected %d got %d", expected, count)
	}
}

func ExampleCharacterCount() {
	count := CharacterCount("string")
	fmt.Println(count)
	// Output: 6
}

func BenchmarkCharacterCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CharacterCount("string")
	}
}
