package pushswap

type Stacks struct {
	A []int
	B []int
}

func NewStacks(numbers []int) Stacks {
	return Stacks{
		A: numbers,
		B: []int{},
	}
}
