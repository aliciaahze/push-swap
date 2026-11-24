package pushswap

import (
	"reflect"
	"testing"
)

func TestLireNombres(t *testing.T) {
	cas := []struct {
		nom       string
		arguments []string
		attendu   []int
		erreur    bool
	}{
		{
			"Nombres valides",
			[]string{"4", "2", "7"},
			[]int{4, 2, 7},
			false,
		},
		{
			"Argument groupe",
			[]string{"4 2 7"},
			[]int{4, 2, 7},
			false,
		},
		{
			"Nombres negatifs",
			[]string{"-5", "0", "12"},
			[]int{-5, 0, 12},
			false,
		},
		{
			"Espaces multiples",
			[]string{"  4   2  7  "},
			[]int{4, 2, 7},
			false,
		},
		{
			"Lettre invalide",
			[]string{"4", "abc", "7"},
			nil,
			true,
		},
		{
			"Nombre decimal",
			[]string{"4", "2.5"},
			nil,
			true,
		},
		{
			"Doublon",
			[]string{"4", "2", "4"},
			nil,
			true,
		},
		{
			"Argument vide",
			[]string{""},
			nil,
			true,
		},
		{
			"Espaces uniquement",
			[]string{"   "},
			nil,
			true,
		},
		{
			"Nombre trop grand",
			[]string{"999999999999999999999999"},
			nil,
			true,
		},
		{
			"Aucun argument",
			[]string{},
			[]int{},
			false,
		},
	}

	for _, test := range cas {
		t.Run(test.nom, func(t *testing.T) {
			resultat, err := LireNombres(test.arguments)

			if (err != nil) != test.erreur {
				t.Fatalf("Erreur inattendue : %v", err)
			}

			if !reflect.DeepEqual(resultat, test.attendu) {
				t.Errorf("Attendu %v, obtenu %v",
					test.attendu, resultat)
			}
		})
	}
}
