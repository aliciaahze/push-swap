package main

import (
	"fmt"
	"os"

	"github.com/aliciaahze/push-swap/internal/pushswap"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	nombres, err := pushswap.LireNombres(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if pushswap.EstTriee(nombres) {
		return
	}

	piles := pushswap.NewStacks(nombres)

	if len(piles.A) == 2 {
		pushswap.SwapA(&piles)
		fmt.Println("sa")
		return
	}

	if len(piles.A) == 3 {
		instructions := pushswap.TrierTrois(&piles)

		for _, instruction := range instructions {
			fmt.Println(instruction)
		}
		return
	}

}
