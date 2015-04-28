package main

import (
	//."driver"
	//."fmt"	
	//."time"
	"control"
	)

func main(){
	
	
	asd := make(chan int)

	control.InitElevator()

	<- asd
	
	return
}
