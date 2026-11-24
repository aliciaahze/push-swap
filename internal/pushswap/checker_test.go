package pushswap

import (
	"reflect"
	"testing"
)

func TestAppliquerInstruction(t *testing.T) {
	piles := NewStacks([]int{3, 1, 2})

	instructions := []string{"sa", "ra"}

	for _, instruction := range instructions {
		if !AppliquerInstruction(&piles, instruction) {
			t.Fatalf("Instruction refusee : %s", instruction)
		}
	}

	attendu := []int{3, 2, 1}

	if !reflect.DeepEqual(piles.A, attendu) {
		t.Errorf("Attendu %v, obtenu %v", attendu, piles.A)
	}
}

func TestInstructionInvalide(t *testing.T) {
	piles := NewStacks([]int{3, 1, 2})

	instructions := []string{"abc", "xx", "", "SA"}

	for _, instruction := range instructions {
		if AppliquerInstruction(&piles, instruction) {
			t.Errorf("Instruction invalide acceptee : %q",
				instruction)
		}
	}

	if !reflect.DeepEqual(piles.A, []int{3, 1, 2}) {
		t.Error("La pile A ne devait pas etre modifiee")
	}

	if len(piles.B) != 0 {
		t.Error("La pile B devait rester vide")
	}
}

func TestToutesLesInstructions(t *testing.T) {
	instructions := []string{
		"sa", "sb", "ss",
		"pa", "pb",
		"ra", "rb", "rr",
		"rra", "rrb", "rrr",
	}

	for _, instruction := range instructions {
		piles := Stacks{
			A: []int{3, 1, 2},
			B: []int{6, 4, 5},
		}

		if !AppliquerInstruction(&piles, instruction) {
			t.Errorf("Instruction valide refusee : %s",
				instruction)
		}
	}
}
