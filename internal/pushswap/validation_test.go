package pushswap

import "testing"

func TestContientDoublon(t *testing.T) {
	nombres := []int{8, 3, 8}

	resultat := ContientDoublon(nombres)

	if resultat != true {
		t.Errorf("Attendu true, obtenu %v", resultat)
	}
}

func TestSansDoublon(t *testing.T) {
	nombres := []int{8, 3, 5}

	resultat := ContientDoublon(nombres)

	if resultat != false {
		t.Errorf("Attendu false, obtenu %v", resultat)
	}
}

func TestSliceVide(t *testing.T) {
	nombres := []int{}

	resultat := ContientDoublon(nombres)

	if resultat != false {
		t.Errorf("Attendu false, obtenu %v", resultat)
	}
}
