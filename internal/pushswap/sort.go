package pushswap

import "sort"

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

func IndiceMinimum(nombres []int) int {
	indiceMin := 0

	for i := 1; i < len(nombres); i++ {
		if nombres[i] < nombres[indiceMin] {
			indiceMin = i
		}
	}

	return indiceMin
}

func TrierQuatreOuCinqNbs(piles *Stacks) []string {
	instructions := []string{}

	if EstTriee(piles.A) {
		return instructions
	}

	for len(piles.A) > 3 {
		indiceMin := IndiceMinimum(piles.A)

		if indiceMin <= len(piles.A)/2 {
			for i := 0; i < indiceMin; i++ {
				RotateA(piles)
				instructions = append(instructions, "ra")
			}
		} else {
			rotations := len(piles.A) - indiceMin

			for i := 0; i < rotations; i++ {
				ReverseRotateA(piles)
				instructions = append(instructions, "rra")
			}
		}

		PushB(piles)
		instructions = append(instructions, "pb")
	}

	instructions = append(instructions, TrierTrois(piles)...)

	for len(piles.B) > 0 {
		PushA(piles)
		instructions = append(instructions, "pa")
	}

	return instructions
}

func ConvertirEnRangs(nombres []int) []int {
	copie := append([]int{}, nombres...)
	sort.Ints(copie)

	rangs := make(map[int]int)

	for i, nombre := range copie {
		rangs[nombre] = i
	}

	resultat := make([]int, len(nombres))

	for i, nombre := range nombres {
		resultat[i] = rangs[nombre]
	}

	return resultat
}

func TrierRadix(piles *Stacks) []string {
	instructions := []string{}

	if EstTriee(piles.A) {
		return instructions
	}

	piles.A = ConvertirEnRangs(piles.A)

	taille := len(piles.A)

	for bit := 0; (taille-1)>>bit > 0; bit++ {
		for i := 0; i < taille; i++ {
			nombre := piles.A[0]

			if (nombre>>bit)&1 == 0 {
				PushB(piles)
				instructions = append(instructions, "pb")
			} else {
				RotateA(piles)
				instructions = append(instructions, "ra")
			}
		}

		for len(piles.B) > 0 {
			PushA(piles)
			instructions = append(instructions, "pa")
		}
	}

	return instructions
}
