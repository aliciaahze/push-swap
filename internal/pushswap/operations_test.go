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

func TestPushAQuandBVide(t *testing.T) {
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

func TestRotateA(t *testing.T) {
	piles := NewStacks([]int{9, 4, 2, 7})

	RotateA(&piles)

	if len(piles.A) != 4 {
		t.Fatalf("La pile A devrait contenir 4 éléments")
	}

	if piles.A[0] != 4 || piles.A[1] != 2 ||
		piles.A[2] != 7 || piles.A[3] != 9 {
		t.Errorf("La pile A devrait contenir [4, 2, 7, 9]")
	}
}

func TestRotateAPileVide(t *testing.T) {
	piles := NewStacks([]int{})

	RotateA(&piles)

	if len(piles.A) != 0 {
		t.Errorf("La pile A devrait rester vide")
	}
}

func TestRotateAUnElement(t *testing.T) {
	piles := NewStacks([]int{7})

	RotateA(&piles)

	if len(piles.A) != 1 {
		t.Fatalf("La pile A devrait contenir un élément")
	}

	if piles.A[0] != 7 {
		t.Errorf("La pile A devrait toujours contenir 7")
	}
}

func TestRotateB(t *testing.T) {
	piles := NewStacks([]int{8, 3})
	piles.B = []int{5, 2, 7}

	RotateB(&piles)

	if len(piles.B) != 3 {
		t.Fatalf("La pile B devrait contenir 3 éléments")
	}

	if piles.B[0] != 2 || piles.B[1] != 7 || piles.B[2] != 5 {
		t.Errorf("La pile B devrait contenir [2, 7, 5]")
	}

	if piles.A[0] != 8 || piles.A[1] != 3 {
		t.Errorf("La pile A ne devrait pas changer")
	}
}

func TestRotateBPileVide(t *testing.T) {
	piles := NewStacks([]int{4, 2})

	RotateB(&piles)

	if len(piles.B) != 0 {
		t.Errorf("La pile B devrait rester vide")
	}
}

func TestRotateBUnElement(t *testing.T) {
	piles := NewStacks([]int{4, 2})
	piles.B = []int{7}

	RotateB(&piles)

	if len(piles.B) != 1 {
		t.Fatalf("La pile B devrait contenir un élément")
	}

	if piles.B[0] != 7 {
		t.Errorf("La pile B devrait toujours contenir 7")
	}
}

func TestRotateBoth(t *testing.T) {
	piles := NewStacks([]int{9, 4, 2})
	piles.B = []int{6, 3, 8}

	RotateBoth(&piles)

	if len(piles.A) != 3 || len(piles.B) != 3 {
		t.Fatalf("Les deux piles devraient contenir 3 éléments")
	}

	if piles.A[0] != 4 || piles.A[1] != 2 || piles.A[2] != 9 {
		t.Errorf("La pile A devrait contenir [4, 2, 9]")
	}

	if piles.B[0] != 3 || piles.B[1] != 8 || piles.B[2] != 6 {
		t.Errorf("La pile B devrait contenir [3, 8, 6]")
	}
}

func TestReverseRotateA(t *testing.T) {
	piles := NewStacks([]int{9, 4, 2, 7})

	ReverseRotateA(&piles)

	if len(piles.A) != 4 {
		t.Fatalf("La pile A devrait contenir 4 éléments")
	}

	if piles.A[0] != 7 || piles.A[1] != 9 ||
		piles.A[2] != 4 || piles.A[3] != 2 {
		t.Errorf("La pile A devrait contenir [7, 9, 4, 2]")
	}
}

func TestReverseRotateAPileVide(t *testing.T) {
	piles := NewStacks([]int{})

	ReverseRotateA(&piles)

	if len(piles.A) != 0 {
		t.Errorf("La pile A devrait rester vide")
	}
}

func TestReverseRotateAUnElement(t *testing.T) {
	piles := NewStacks([]int{7})

	ReverseRotateA(&piles)

	if len(piles.A) != 1 {
		t.Fatalf("La pile A devrait contenir un élément")
	}

	if piles.A[0] != 7 {
		t.Errorf("La pile A devrait toujours contenir 7")
	}
}

func TestReverseRotateB(t *testing.T) {
	piles := NewStacks([]int{8, 4})
	piles.B = []int{9, 2, 6}

	ReverseRotateB(&piles)

	if len(piles.B) != 3 {
		t.Fatalf("La pile B devrait contenir 3 éléments")
	}

	if piles.B[0] != 6 || piles.B[1] != 9 || piles.B[2] != 2 {
		t.Errorf("La pile B devrait contenir [6, 9, 2]")
	}

	if piles.A[0] != 8 || piles.A[1] != 4 {
		t.Errorf("La pile A ne devrait pas changer")
	}
}

func TestReverseRotateBPileVide(t *testing.T) {
	piles := NewStacks([]int{4, 2})

	ReverseRotateB(&piles)

	if len(piles.B) != 0 {
		t.Errorf("La pile B devrait rester vide")
	}
}

func TestReverseRotateBUnElement(t *testing.T) {
	piles := NewStacks([]int{4, 2})
	piles.B = []int{7}

	ReverseRotateB(&piles)

	if len(piles.B) != 1 {
		t.Fatalf("La pile B devrait contenir un élément")
	}

	if piles.B[0] != 7 {
		t.Errorf("La pile B devrait toujours contenir 7")
	}
}

func TestReverseRotateBoth(t *testing.T) {
	piles := NewStacks([]int{9, 4, 2})
	piles.B = []int{6, 3, 8}

	ReverseRotateBoth(&piles)

	if len(piles.A) != 3 || len(piles.B) != 3 {
		t.Fatalf("Les deux piles devraient contenir 3 éléments")
	}

	if piles.A[0] != 2 || piles.A[1] != 9 || piles.A[2] != 4 {
		t.Errorf("La pile A devrait contenir [2, 9, 4]")
	}

	if piles.B[0] != 8 || piles.B[1] != 6 || piles.B[2] != 3 {
		t.Errorf("La pile B devrait contenir [8, 6, 3]")
	}
}
