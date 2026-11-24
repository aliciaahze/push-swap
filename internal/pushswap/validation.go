package pushswap

func ContientDoublon(nombres []int) bool {
	vus := make(map[int]bool)

	for _, nombre := range nombres {
		if vus[nombre] {
			return true
		}

		vus[nombre] = true
	}

	return false
}
