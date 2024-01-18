package arrays_slice

func Sum(numbers [5]int) (sum int) {

	for _, number := range numbers {
		sum += number
	}

	return
}

func SumSlice(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func AllSum(numbers ...[]int) (sum []int) {

	for _, number := range numbers {
		sum = append(sum, SumSlice(number))
	}

	return
}

// quando usamos [1:] - estamos dizendo: pegue da posição 1 até o final.
func SumAllRest(numbers ...[]int) (sum []int) {
	for _, number := range numbers {

		if len(number) == 0 {
			sum = append(sum, 0)
			continue
		}

		final := number[1:]
		sum = append(sum, SumSlice(final))
	}

	return
}
