package arrays_slice

import (
	"reflect"
	"testing"
)

/*
		O Go não deixa usar operadores de igualdade com slices. Para patricidade, podemos usar o reflect.DeepEqual
	 	Ele é util para verificar se dyas variáveis são iguais
		É importante saber que o reflect.DeepEqual não tem segurança de tipos.
		Ou seja, se o código espera int e eu passo string, ele irá copilar
*/
func TestSum(t *testing.T) {
	got := Sum([5]int{2, 4, 5, 6, 7})
	expected := 24

	if got != expected {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestSumSlice(t *testing.T) {
	got := SumSlice([]int{2, 4, 6, 8, 9})
	expected := 29

	if got != expected {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestAllSum(t *testing.T) {
	got := AllSum([]int{0, 3}, []int{1, 6})
	expected := []int{3, 7}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected %v but got %v", expected, got)
	}
}

func TestSumAllRest(t *testing.T) {

	verifySum := func(t *testing.T, got, expected []int) {
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %v but got %v", expected, got)
		}
	}

	t.Run("sum all rest slice", func(t *testing.T) {
		got := SumAllRest([]int{2, 3}, []int{1, 6})
		expected := []int{3, 6}

		verifySum(t, got, expected)
	})

	t.Run("sum all slice empty", func(t *testing.T) {
		got := SumAllRest([]int{}, []int{0, 3})
		expected := []int{0, 3}

		verifySum(t, got, expected)
	})
}
