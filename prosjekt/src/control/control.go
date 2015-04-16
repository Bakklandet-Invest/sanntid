package control

import (
	"network"
	
	)

type Control struct {
	elevators map[int]*network.Elevator
	orders map[int][2]bool
	connected map[string]bool
}

func (cont *Control) InitMaps() {
	
	cont.elevators = make(map[string]*network.Elevator)
	cont.connected = make(map[string]bool)
	cont.orders = make(map[int][2]bool)
	
	return
}


