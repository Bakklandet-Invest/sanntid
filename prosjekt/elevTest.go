package main

import (
	//."driver"
	//."fmt"	
	//."time"
	"control"
	)

func main(){
	go slave
}
	
func slave() {
	
	
	asd := make(chan int)

	updateChan := make(msg Message)
	exOrderChan := make(order ButtonSignal)
	fromMaster := make(order ButtonSignal)
	readUDPButtonSig := make(order ButtonSignal)

	go control.InitElevator(updateChan, exOrderChan, fromMaster)
	go 
		
	koble til bla bla bla
	
	if (myID = masterID)
		go master
		return
	for{
		select{
			case update := <-updateChan
				send update

			case exOrder := <- exOrderChan
				send ordre til master
				if !masterAlive
					sjekke kartet, finne hvem som er master
					go master
					return
	
	

	<- asd
	
	return
}
