package strings

func LCS(s1, s2 string) string {

	sra := []rune(s1)
	sla := len(sra)

	srb := []rune(s2)
	slb := len(srb)

	if sla == 0 && slb == 0 {
		return ""
	}

	table := make([][]int, sla+1)
	for i := 0; i < sla+1; i++ {
		table[i] = make([]int, slb+1)
	}

	for i := 0; i < sla; i++ {
		for j := 0; j < slb; j++ {
			if sra[i] == srb[j] {
				table[i+1][j+1] = table[i][j] + 1

			} else {
				if table[i+1][j] > table[i][j+1] {
					table[i+1][j+1] = table[i+1][j]
				} else {
					table[i+1][j+1] = table[i][j+1]
				}
			}
		}
	}

	size := table[sla][slb]
	result := make([]rune, 0, size)
	for i, j := sla, slb; i > 0 && j > 0; {
		if table[i][j] == table[i-1][j] {
			i--
		} else if table[i][j] == table[i][j-1] {
			j--
		} else {
			result = append(result, sra[i-1])
			i--
			j--
		}

	}

	for i := 0; i < size/2; i++ {
		result[i], result[size-i-1] = result[size-i-1], result[i]
	}

	return string(result)
}
