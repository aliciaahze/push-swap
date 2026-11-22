package pushswap

func EstTriee(nombres []int) bool {
	for i := 0; i < len(nombres)-1; i++ {
		if nombres[i] > nombres[i+1] {
			return false
		}
	}

	return true
}
