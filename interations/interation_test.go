package interations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	expected := "aaaaa"

	if reflect.DeepEqual(expected, got) {
		t.Errorf("expected %s but got %s", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a"))
	//Output: aaaaa
}
