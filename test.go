package main

import(
	."fmt"
	//."math"
	)

type Matrix [4][3]int

type rompe struct {
	matrise Matrix
}
	
func main(){
	hei := new(rompe)
	Println(hei.matrise)
	var matr Matrix
	matr[1][1] = 5
	Println(matr)

	for i := 0; i <= 5; i++ {
		Println(i)
	}

	Println("---------")
	asd := -matr[1][1]
	Println(asd)
}


