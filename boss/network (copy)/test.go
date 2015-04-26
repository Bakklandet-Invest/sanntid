package main

import(
	"fmt"
	)
	
func main(){
	g := f(3)
	fmt.Println(g)
}

func f(int int)int{
	g := 2*int
	return g
}
