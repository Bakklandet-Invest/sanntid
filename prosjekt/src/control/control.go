package control
/*
import (
	"network"
	"elevator"
	)

func slave() {
	updateChan := make(chan Message)
	exOrderChan := make(chan ButtonSignal)
	fromMaster := make(chan ButtonSignal)
	UDPButtonSig := make(chan ButtonSignal)
	UpdateMsg := make(chan Message)

	go control.InitElevator(updateChan, exOrderChan, fromMaster)
	go messageHandler(UDPButtonSig, UpdateMsg) 
	
	
	
	koble til bla bla bla
	
	
	if (myID = masterID)
		go master
		return
	for{
		select{
			case msg := <-MessageChan
				MessageHandler(msg)

			case exOrder := <- exOrderChan
				send ordre til master
				if !masterAlive
					sjekke kartet, finne hvem som er master
					go master
					return
	
	




	

	return
}
*/
