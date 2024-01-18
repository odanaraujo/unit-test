package interations

const qtdeRepeat = 5

func Repeat(caracterer string) string {
	var caractererRepeat string
	for i := 0; i < qtdeRepeat; i++ {
		caractererRepeat += caracterer
	}

	return caracterer
}
