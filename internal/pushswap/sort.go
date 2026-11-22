package pushswap

func EstTriee(nombres []int) bool {
	for i := 0; i < len(nombres)-1; i++ {
		if nombres[i] > nombres[i+1] {
			return false
		}
	}

	return true
}

func TrierTrois(piles *Stacks) []string {
	instructions := []string{}

	if len(piles.A) != 3 || EstTriee(piles.A) {
		return instructions
	}

	a := piles.A[0]
	b := piles.A[1]
	c := piles.A[2]

	if a > b && b < c && a < c {
		SwapA(piles)
		instructions = append(instructions, "sa")
	} else if a > b && b > c {
		SwapA(piles)
		ReverseRotateA(piles)
		instructions = append(instructions, "sa", "rra")
	} else if a > b && b < c && a > c {
		RotateA(piles)
		instructions = append(instructions, "ra")
	} else if a < b && b > c && a < c {
		SwapA(piles)
		RotateA(piles)
		instructions = append(instructions, "sa", "ra")
	} else if a < b && b > c && a > c {
		ReverseRotateA(piles)
		instructions = append(instructions, "rra")
	}

	return instructions
}
