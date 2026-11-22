package pushswap

import (
	"errors"
	"strconv"
)

func LireNombres(arguments []string) ([]int, error) {
	nombres := []int{}

	for _, texte := range arguments {
		nombre, err := strconv.Atoi(texte)
		if err != nil {
			return nil, errors.New("argument invalide")
		}

		nombres = append(nombres, nombre)
	}

	if ContientDoublon(nombres) {
		return nil, errors.New("doublon détecté")
	}

	return nombres, nil
}
