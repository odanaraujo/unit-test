package integers

import "math"

func Sum(x, y int) int {
	return int(math.Abs(float64(x + y)))
}
