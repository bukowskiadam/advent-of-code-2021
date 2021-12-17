package utils

func FindMinInt(input []int) int {
	m := 0
	for i, e := range input {
		if i == 0 || e < m {
			m = e
		}
	}
	return m
}

func FindMaxInt(input []int) int {
	m := 0
	for i, e := range input {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}
