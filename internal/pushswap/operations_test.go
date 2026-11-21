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

func TestSwapB(t *testing.T) {
	piles := NewStacks([]int{10, 4})

	piles.B = []int{6, 2, 9}

	SwapB(&piles)

	if piles.B[0] != 2 {
		t.Errorf("Le premier élément de B devrait être 2")
	}

	if piles.B[1] != 6 {
		t.Errorf("Le deuxième élément de B devrait être 6")
	}

	if piles.B[2] != 9 {
		t.Errorf("Le troisième élément de B devrait rester 9")
	}
}

func TestSwapBoth(t *testing.T) {
	piles := NewStacks([]int{8, 3, 5})
	piles.B = []int{7, 2, 1}

	SwapBoth(&piles)

	if piles.A[0] != 3 || piles.A[1] != 8 {
		t.Errorf("La pile A n'a pas été correctement échangée")
	}

	if piles.B[0] != 2 || piles.B[1] != 7 {
		t.Errorf("La pile B n'a pas été correctement échangée")
	}
}

func TestPushB(t *testing.T) {
	piles := NewStacks([]int{9, 4, 2})
	piles.B = []int{6, 1}
	PushB(&piles)
	if len(piles.A) != 2 {
		t.Errorf("La pile A devrait contenir 2 éléments")
	}

	if len(piles.B) != 3 {
		t.Errorf("La pile B devrait contenir 3 éléments")
	}
	if piles.A[0] != 4 || piles.A[1] != 2 {
		t.Errorf("La pile A devrait contenir [4, 2]")
	}

	if piles.B[0] != 9 || piles.B[1] != 6 || piles.B[2] != 1 {
		t.Errorf("La pile B devrait contenir [9, 6, 1]")
	}
}

func TestPushBPileAVide(t *testing.T) {
	piles := NewStacks([]int{})
	piles.B = []int{7, 2}

	PushB(&piles)

	if len(piles.A) != 0 {
		t.Errorf("La pile A devrait rester vide")
	}

	if len(piles.B) != 2 {
		t.Errorf("La pile B devrait conserver 2 éléments")
	}

	if piles.B[0] != 7 || piles.B[1] != 2 {
		t.Errorf("La pile B ne devrait pas changer")
	}
}

func TestPushA(t *testing.T) {
	piles := NewStacks([]int{4, 2})
	piles.B = []int{7, 9, 3}

	PushA(&piles)

	if len(piles.A) != 3 {
		t.Errorf("La pile A devrait contenir 3 éléments")
	}

	if len(piles.B) != 2 {
		t.Errorf("La pile B devrait contenir 2 éléments")
	}

	if len(piles.A) == 3 {
		if piles.A[0] != 7 || piles.A[1] != 4 || piles.A[2] != 2 {
			t.Errorf("La pile A devrait contenir [7, 4, 2]")
		}
	}

	if len(piles.B) == 2 {
		if piles.B[0] != 9 || piles.B[1] != 3 {
			t.Errorf("La pile B devrait contenir [9, 3]")
		}
	}
}

func TestPushBPileVide(t *testing.T) {
	piles := NewStacks([]int{4, 2})

	PushA(&piles)

	if len(piles.A) != 2 {
		t.Errorf("La pile A devrait contenir 2 éléments")
	}

	if piles.A[0] != 4 || piles.A[1] != 2 {
		t.Errorf("La pile A ne devrait pas changer")
	}

	if len(piles.B) != 0 {
		t.Errorf("La pile B devrait rester vide")
	}
}
