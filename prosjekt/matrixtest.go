package main

import(
	"fmt"
)

const NumButtons = 3
const NumFloors = 4

type Matrix [4][3]bool

func main(){
	//hengekanal := make(chan int)
	var m1 Matrix
	var m2 Matrix
	m1[2][2] = true
	
	m3 := orMatrixCompare(m1, m2)
	fmt.Println(m3)
	if m1 == m2{
		fmt.Println("yes!")
		fmt.Println(m1[2][2])
	} else{
		fmt.Println("nope")
		fmt.Println(m1[2][2])
	}
	
	//<-hengekanal
}

func orMatrixCompare(m1 Matrix, m2 Matrix) (Matrix){
	var m3 Matrix
	for i := 0; i < NumFloors; i++ {
		for j := 0; j < NumButtons; j++{
			if m1[i][j] == true || m2[i][j] == true{
				m3[i][j] = true
			}
		}
	}
	return m3
}
