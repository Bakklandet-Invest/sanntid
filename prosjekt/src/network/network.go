package network

import (
	"net"
	//"time"
	//"encoding/json"
	"github.com/fredosavage/util"
)

type Message struct{
	elev *Elevator
	takenOrder bool
	floorOrder
}

type Elevator struct {
	stopList *list.LinkedList
	speed int
	direction int
	currentFloor int
	out chan string
}

//func 