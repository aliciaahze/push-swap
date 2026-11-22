package pushswap

import "testing"

func TestTrierRadix(t *testing.T) {
	cas := []struct {
		nom     string
		nombres []int
	}{
		{"Quatre nombres", []int{4, 3, 2, 1}},
		{"Nombres negatifs", []int{80, -10, 35, 200}},
		{"Cinq nombres", []int{5, 1, 4, 2, 3}},
		{"Dix nombres", []int{9, 2, 7, 1, 5, 0, 8, 3, 6, 4}},
		{"Deja triee", []int{1, 2, 3, 4}},
	}

	for _, test := range cas {
		t.Run(test.nom, func(t *testing.T) {
			piles := NewStacks(test.nombres)

			TrierRadix(&piles)

			if !EstTriee(piles.A) {
				t.Errorf("Pile A non triee : %v", piles.A)
			}

			if len(piles.B) != 0 {
				t.Errorf("Pile B non vide : %v", piles.B)
			}
		})
	}
}
