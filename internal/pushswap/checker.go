package pushswap

func AppliquerInstruction(piles *Stacks, instruction string) bool {
	switch instruction {
	case "sa":
		SwapA(piles)
	case "sb":
		SwapB(piles)
	case "ss":
		SwapBoth(piles)
	case "pa":
		PushA(piles)
	case "pb":
		PushB(piles)
	case "ra":
		RotateA(piles)
	case "rb":
		RotateB(piles)
	case "rr":
		RotateBoth(piles)
	case "rra":
		ReverseRotateA(piles)
	case "rrb":
		ReverseRotateB(piles)
	case "rrr":
		ReverseRotateBoth(piles)
	default:
		return false
	}

	return true
}
