package utils

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Factorial(x int) int {
	return x * (x + 1) / 2
}
