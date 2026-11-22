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

	fmt.Println("Pile A :", piles.A)
	fmt.Println("Pile B :", piles.B)
}
