package control

import (
	//"llist"
	"strconv"
	"net"
	."driver"
	."fmt"
	."time" 
	)

type Matrix [4][3]bool

type Elevator struct {
	id int
	orderMatrix Matrix
	direction int
	currentFloor int
	destination int

	
	// newOrderNotifyChan chan bool 
		// bruke en kanal for å gi beskjed når en ny ordre kommer for å unngå
		// forløkker som kjører konstant. En input i notify setter i gang 
		// GetNextDestination som igjen setter i gang en case i Run
}

func FindElevID() int {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                IDstr := ipnet.IP.String()
                IDstr = IDstr[12:]
                ID, _ := strconv.Atoi(IDstr)
                return ID
            }
        }
    }
    return 0
}

func InitElevator() *Elevator {

	if Elev_init() == 0 {
		Println("Could not initialize elevator")
	}
		
	Elev_set_stop_lamp(1)
	
	e := new(Elevator)
	e.id = FindElevID()

	//e.stopList = llist.New()
	
	if Elev_get_floor_sensor_signal() == -1 {
		Elev_set_speed(-300)
		for Elev_get_floor_sensor_signal() == -1 {}
	}
	Elev_set_speed(0)
	e.currentFloor = Elev_get_floor_sensor_signal()

	/*
	updateChan chan Elevator
	intOrderChan chan ButtonSignal
	extOrderChan chan ButtonSignal
	nextDestinationChan chan int	
	*/

	//updateChan := make(chan Elevator, 1)
	intOrderChan := make(chan ButtonSignal, 1)
	extOrderChan := make(chan ButtonSignal, 1)
	masterOrderChan := make(chan ButtonSignal, 1)
	//nextDestinationChan := make(chan int, 1)

	go e.OrderHandler(intOrderChan, extOrderChan)
	//go Elev_get_order(intOrderChan, extOrderChan)
	go e.Run()

	return e
}


func (e *Elevator) OrderHandler(intOrderChan chan ButtonSignal, extOrderChan chan ButtonSignal) {
	// external order, send til master
	// internal order legg til først i stopplisten 
	// (kun etter orders som er i etasjer over og i riktig retning)
	var newOrder ButtonSignal
	for{
		select{
			case newOrder = <- extOrderChan:
				// send newOrder til master
			case newOrder = <- intOrderChan:
				// ..
		}
	}
}

func (e *Elevator) Run() {
	// go funksjonen som skriver til nextDestinationChan
	// go funksjonen som oppdaterer elevator variablene
	for{
		/* if (Elev_get_obstruction_signal() == 1) {
				handle obstruction 
			}
		else { */
				if Elev_get_stop_signal() {
					Elev_set_speed(0)
					Elev_set_stop_lamp(-1)
					// hva skal skje når stop-knappen trykkes?
				} else if (e.destination < e.currentFloor) {
					e.direction = -1
					Elev_set_speed(-300)
				} else if (e.destination > e.currentFloor) {
					e.direction = 1
					Elev_set_speed(300)
				} else {

					Elev_set_speed(0)

					// e.stopList.Remove(l.Front()) //fjern oppfylt ordre fra køen
					Elev_set_door_open_lamp(1)

				Sleep(Second)
				Elev_set_door_open_lamp(0)
			}




	}	//for
}	//func

// fiks så den bruker chanal og funker for vårt oppsett


func (e *Elevator) addOrder(intOrderChan chan ButtonSignal, masterOrderChan chan ButtonSignal) {
	var order ButtonSignal	
	for{
		select{
			case order = <- masterOrderChan:
				e.orderMatrix[order.Floor][order.Button - 1] = true
			case order = <- intOrderChan:
				e.orderMatrix[order.Floor][order.Button - 1] = true
		}
	}
}

func (e *Elevator) removeOrder(floor int, button int) {
	e.orderMatrix[order.Floor][order.Button - 1] = false
}


func (e *Elevator) getNewDirection() {
		// return int eller send på channel ?
	
	if (e.orderMatrix[e.currentFloor][0] || e.orderMatrix[e.currentFloor][1] || e.orderMatrix[e.currentFloor][2]) {
			 // return eller send
	}
	dist := N_FLOORS
	next := e.currentFloor
	for i := 0; i <= 3; i++ {
		if e.orderMatrix[i][2] {
			if math.Abs(e.currentFloor - i) < dist {
				dist = math.Abs(e.currentFloor - i)
				next = i
			}
		}
	}	
	if next != e.currentFloor {
		// return eller send
	}
	
	for i := 0; i <= 3; i++ { // finn nederste som vil opp
		if orderMatrix[i][0] {
			dist = math.Abs(e.currentFloor - i)
			next = i
		}	
	}	
	for i := 3; i >= 0; i-- { // finn øverste ordre som vil ned
		if orderMatrix[i][1] {
			if math.Abs(e.currentFloor - i) < dist {
				dist = math.Abs(e.currentFloor - i)
				next = i
			}
		}
	}
	// return eller send next
}


func (e *Elevator) orderInOtherDir() bool {
	if e.direction > 0 {
		for i := 0; i <= 3; i++ {
			if orderMatrix[i][1] {
				return true
			}
		}
	} else if e.direction < 0 {
		for i := 0; i <= 3; i++ {
			if orderMatrix[i][0] {
				return true
			}		
		}
	}
}

func (e *Elevator) orderOnCurrentFloor() bool {	
	if (e.orderMatrix[Elev_get_floor_sensor_signal()][0] || e.orderMatrix[Elev_get_floor_sensor_signal()][1] || e.orderMatrix[Elev_get_floor_sensor_signal()][2]) {
		return true
	}
	return false
}

func (e *Elevator) orderInCurrentDir() bool {
	if e.direction > 0 {
		for i := e.currentFloor; i <= 3; i++ {
			if orderMatrix[i][0] || orderMatrix[i][2] {
				return true
			}
		}	
	}	
	else if e.direction < 0
		for i := e.currentFloor; i >= 0; i-- {
			if orderMatrix[i][1] || orderMatrix[i][2] {
				return true
			}
		}
	}
	return false
}
