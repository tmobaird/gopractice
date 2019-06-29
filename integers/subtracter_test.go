package integers

import (
	"fmt"
	"testing"
)

func TestSubtracter(t *testing.T) {
	difference := Subtract(5, 3)
	expected := 2

	if difference != expected {
		t.Errorf("expected '%d' but got '%d'", expected, difference)
	}
}

func ExampleSubtract() {
	difference := Subtract(5, 1)
	fmt.Println(difference)
	// Output: 4
}
