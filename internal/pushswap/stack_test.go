package pushswap

import "testing"

func TestNewStacks(t *testing.T) {
	nombres := []int{4, 2, 7}

	piles := NewStacks(nombres)

	if len(piles.A) != 3 {
		t.Errorf("La pile A devrait contenir 3 éléments")
	}

	if len(piles.B) != 0 {
		t.Errorf("La pile B devrait être vide")
	}

	for i, nombre := range nombres {
		if piles.A[i] != nombre {
			t.Errorf("Valeur incorrecte à l'indice %d", i)
		}
	}

}
