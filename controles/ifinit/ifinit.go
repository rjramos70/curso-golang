package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)

}

func main() {
	if i := numeroAleatorio(); i > 5 { // também suportado no Switch
		fmt.Printf("[%v] Ganhou!!!", i)
	} else {
		fmt.Printf("[%v] Perdeu!!!", i)
	}
}
