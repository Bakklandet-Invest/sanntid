package main

import(
	."fmt"
	."math"
	)


	
func main(){
	type Matrix [4][3]int
	var matr Matrix
	matr[1][1] = 5
	Println(matr)

	for i := 0; i <= 5; i++ {
		Println(i)
	}

	Println(Abs(-1))
}


