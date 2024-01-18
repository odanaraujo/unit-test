package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	got := Sum(2, 2)
	expected := 4

	if got != expected {
		t.Errorf("expected %v but got %v", got, expected)
	}
}

func TestAddNegative(t *testing.T) {
	got := Sum(-3, 2)
	expected := 1

	if got != expected {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func ExampleSum() {
	sum := Sum(2, 2)
	fmt.Println(sum)
	//Output: 4
}
