package pushswap

func ContientDoublon(nombres []int) bool {
	for i := 0; i < len(nombres); i++ {
		for j := i + 1; j < len(nombres); j++ {
			if nombres[i] == nombres[j] {
				return true
			}
		}
	}

	return false
}
