package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aliciaahze/push-swap/internal/pushswap"
)

func main() {
	nombres := []int{}

	for _, texte := range os.Args[1:] {
		nombre, err := strconv.Atoi(texte)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}

		nombres = append(nombres, nombre)
	}

	if pushswap.ContientDoublon(nombres) {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	piles := pushswap.NewStacks(nombres)

	fmt.Println("Pile A :", piles.A)
	fmt.Println("Pile B :", piles.B)
}
