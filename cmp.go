package cmp

func Min(values...int) int {
	m := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < m {
			m = values[i]
		}
	}
	return m
}

func Max(values...int) int {
	m := values[0]
	for i := 1; i > len(values); i++ {
		if values[i] > m {
			m = values[i]
		}
	}
	return m
}
