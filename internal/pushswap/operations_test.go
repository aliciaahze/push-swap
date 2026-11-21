package pushswap

import "testing"

func TestSwapA(t *testing.T) {
	piles := NewStacks([]int{10, 4, 7, 2})

	SwapA(&piles)

	if piles.A[0] != 4 {
		t.Errorf("Le premier élément devrait être 4")
	}

	if piles.A[1] != 10 {
		t.Errorf("Le deuxième élément devrait être 10")
	}
}
func TestSwapAUnElement(t *testing.T) {
	piles := NewStacks([]int{8})

	SwapA(&piles)

	if len(piles.A) != 1 || piles.A[0] != 8 {
		t.Errorf("La pile A ne devrait pas changer")
	}
}

func TestSwapAPileVide(t *testing.T) {
	piles := NewStacks([]int{})

	SwapA(&piles)

	if len(piles.A) != 0 {
		t.Errorf("La pile A devrait rester vide")
	}
}
