package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Número de processadores físicos: %d", runtime.NumCPU())
	// Número de processadores físicos: 4
}
