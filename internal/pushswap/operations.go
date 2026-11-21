package pushswap

func SwapA(piles *Stacks) {
	if len(piles.A) < 2 {
		return
	}

	piles.A[0], piles.A[1] = piles.A[1], piles.A[0]
}

func SwapB(piles *Stacks) {
	if len(piles.B) < 2 {
		return
	}

	piles.B[0], piles.B[1] = piles.B[1], piles.B[0]
}

func SwapBoth(piles *Stacks) {
	SwapA(piles)
	SwapB(piles)
}

func PushB(piles *Stacks) {
	if len(piles.A) == 0 {
		return
	}

	premier := piles.A[0]
	piles.A = piles.A[1:]
	piles.B = append([]int{premier}, piles.B...)
}

func PushA(piles *Stacks) {
	if len(piles.B) == 0 {
		return
	}

	premier := piles.B[0]

	piles.B = piles.B[1:]

	piles.A = append([]int{premier}, piles.A...)
}

func RotateA(piles *Stacks) {
	if len(piles.A) < 2 {
		return
	}

	premier := piles.A[0]
	piles.A = piles.A[1:]
	piles.A = append(piles.A, premier)
}

func RotateB(piles *Stacks) {
	if len(piles.B) < 2 {
		return
	}

	premier := piles.B[0]
	piles.B = piles.B[1:]
	piles.B = append(piles.B, premier)
}

func RotateBoth(piles *Stacks) {
	RotateA(piles)
	RotateB(piles)
}

func ReverseRotateA(piles *Stacks) {
	if len(piles.A) < 2 {
		return
	}

	dernier := piles.A[len(piles.A)-1]
	piles.A = piles.A[:len(piles.A)-1]
	piles.A = append([]int{dernier}, piles.A...)
}

func ReverseRotateB(piles *Stacks) {
	if len(piles.B) < 2 {
		return
	}

	dernier := piles.B[len(piles.B)-1]
	piles.B = piles.B[:len(piles.B)-1]
	piles.B = append([]int{dernier}, piles.B...)
}

func ReverseRotateBoth(piles *Stacks) {
	ReverseRotateA(piles)
	ReverseRotateB(piles)
}
