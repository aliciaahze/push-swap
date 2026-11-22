package pushswap

import (
	"errors"
	"strconv"
	"strings"
)

func LireNombres(arguments []string) ([]int, error) {
	nombres := []int{}

	for _, argument := range arguments {
		morceaux := strings.Fields(argument)

		if len(morceaux) == 0 {
			return nil, errors.New("argument vide")
		}

		for _, texte := range morceaux {
			nombre, err := strconv.Atoi(texte)
			if err != nil {
				return nil, errors.New("argument invalide")
			}

			nombres = append(nombres, nombre)
		}
	}

	if ContientDoublon(nombres) {
		return nil, errors.New("doublon détecté")
	}

	return nombres, nil
}
