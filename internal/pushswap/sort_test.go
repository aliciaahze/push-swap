package pushswap

import (
	"reflect"
	"testing"
)

func TestTrierTrois(t *testing.T) {
	cas := [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	}

	for _, nombres := range cas {
		piles := NewStacks(append([]int{}, nombres...))

		instructions := TrierTrois(&piles)

		if !reflect.DeepEqual(piles.A, []int{1, 2, 3}) {
			t.Errorf("Tri incorrect pour %v", nombres)
		}

		if len(piles.B) != 0 {
			t.Error("La pile B devait rester vide")
		}

		if len(instructions) > 2 {
			t.Errorf("Trop d'instructions pour %v", nombres)
		}

		verification := NewStacks(
			append([]int{}, nombres...),
		)

		for _, instruction := range instructions {
			if !AppliquerInstruction(&verification, instruction) {
				t.Fatalf("Instruction invalide : %s", instruction)
			}
		}

		if !reflect.DeepEqual(verification.A, []int{1, 2, 3}) ||
			len(verification.B) != 0 {
			t.Errorf("Instructions incorrectes pour %v", nombres)
		}
	}
}

func TestTrierQuatreOuCinqNbs(t *testing.T) {
	cas := [][]int{
		{4, 3, 2, 1},
		{4, 1, 5, 2, 3},
		{5, 4, 3, 2, 1},
		{80, -10, 35, 200, -5},
		{1, 2, 3, 4, 5},
	}

	for _, nombres := range cas {
		piles := NewStacks(append([]int{}, nombres...))
		instructions := TrierQuatreOuCinqNbs(&piles)

		if !EstTriee(piles.A) {
			t.Errorf("Pile A non triee : %v", piles.A)
		}

		if len(piles.B) != 0 {
			t.Errorf("Pile B non vide : %v", piles.B)
		}

		verification := NewStacks(append([]int{}, nombres...))

		for _, instruction := range instructions {
			if !AppliquerInstruction(&verification, instruction) {
				t.Fatalf("Instruction invalide : %s", instruction)
			}
		}

		if !EstTriee(verification.A) || len(verification.B) != 0 {
			t.Errorf("Les instructions ne trient pas %v", nombres)
		}
	}
}

func TestIndiceMinimum(t *testing.T) {
	cas := []struct {
		nombres []int
		attendu int
	}{
		{[]int{1, 5, 3}, 0},
		{[]int{8, 4, 9, 2, 6}, 3},
		{[]int{4, 5, 1}, 2},
		{[]int{4, -8, 2}, 1},
		{[]int{7}, 0},
	}

	for _, test := range cas {
		resultat := IndiceMinimum(test.nombres)

		if resultat != test.attendu {
			t.Errorf("Pour %v : indice attendu %d, obtenu %d",
				test.nombres, test.attendu, resultat)
		}
	}
}

func TestTrierRadix(t *testing.T) {
	cas := []struct {
		nom     string
		nombres []int
		attendu []int
	}{
		{
			"Quatre nombres",
			[]int{4, 3, 2, 1},
			[]int{1, 2, 3, 4},
		},
		{
			"Nombres negatifs",
			[]int{80, -10, 35, 200},
			[]int{-10, 35, 80, 200},
		},
		{
			"Cinq nombres",
			[]int{5, 1, 4, 2, 3},
			[]int{1, 2, 3, 4, 5},
		},
		{
			"Dix nombres",
			[]int{9, 2, 7, 1, 5, 0, 8, 3, 6, 4},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"Deja triee",
			[]int{1, 2, 3, 4},
			[]int{1, 2, 3, 4},
		},
	}

	for _, test := range cas {
		t.Run(test.nom, func(t *testing.T) {
			piles := NewStacks(
				append([]int{}, test.nombres...),
			)

			instructions := TrierRadix(&piles)

			if !EstTriee(piles.A) {
				t.Errorf("Pile A non triee : %v", piles.A)
			}

			if len(piles.B) != 0 {
				t.Error("Pile B non vide")
			}

			verification := NewStacks(
				append([]int{}, test.nombres...),
			)

			for _, instruction := range instructions {
				if !AppliquerInstruction(&verification, instruction) {
					t.Fatalf("Instruction invalide : %s",
						instruction)
				}
			}

			if !reflect.DeepEqual(verification.A, test.attendu) {
				t.Errorf("Attendu %v, obtenu %v",
					test.attendu, verification.A)
			}

			if len(verification.B) != 0 {
				t.Error("La pile B du checker devait etre vide")
			}

			if EstTriee(test.nombres) && len(instructions) != 0 {
				t.Error("Une pile triee ne necessite aucune instruction")
			}
		})
	}
}

func TestEstTriee(t *testing.T) {
	cas := []struct {
		nombres []int
		attendu bool
	}{
		{[]int{1, 2, 3}, true},
		{[]int{3, 2, 1}, false},
		{[]int{1, 3, 2}, false},
		{[]int{-5, -2, 0, 4}, true},
		{[]int{7}, true},
		{[]int{}, true},
	}

	for _, test := range cas {
		resultat := EstTriee(test.nombres)

		if resultat != test.attendu {
			t.Errorf("Pour %v : attendu %v, obtenu %v",
				test.nombres, test.attendu, resultat)
		}
	}
}

func TestConvertirEnRangs(t *testing.T) {
	nombres := []int{80, -10, 35, 200}

	attendu := []int{2, 0, 1, 3}

	resultat := ConvertirEnRangs(nombres)

	if !reflect.DeepEqual(resultat, attendu) {
		t.Errorf("Attendu %v, obtenu %v", attendu, resultat)
	}

	if !reflect.DeepEqual(nombres, []int{80, -10, 35, 200}) {
		t.Error("Les nombres d'origine ont ete modifies")
	}
}
