package main

import (
	"fmt"
	"reflect"
	"time"
)

func tipo(i interface{}) string {
	switch v := i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return fmt.Sprintf("não sei [tipo %s] ", reflect.TypeOf(v))
		// return reflect.TypeOf(v).String()
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
