package main

import (
	"fmt"
	)

func main() {
	balle := "balle"
	bal := balle[:3]
	le := balle[3:]

	fmt.Println(bal + "   " + le)


	tall := "1234567890101112"
	str := []byte(tall)
	recAddr := string(str)[:9]
	fmt.Println(recAddr)
	return
}
