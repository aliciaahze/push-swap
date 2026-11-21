package pushswap

func SwapA(piles *Stacks) {
	if len(piles.A) < 2 {
		return
	}

	piles.A[0], piles.A[1] = piles.A[1], piles.A[0]
}
