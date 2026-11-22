package main

import (
	"bufio"
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

	piles := pushswap.NewStacks(nombres)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		instruction := scanner.Text()

		if !pushswap.AppliquerInstruction(&piles, instruction) {
			fmt.Fprintln(os.Stderr, "Error")
			os.Exit(1)
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	if pushswap.EstTriee(piles.A) && len(piles.B) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
