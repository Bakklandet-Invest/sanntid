package control

import (
	//"llist"
	"strconv"
	"net"
	."driver"
	."fmt"
	."time" 
	"math"
	)

type Matrix [4][3]bool

type Elevator struct {
	id int
	orderMatrix Matrix
	direction int
	currentFloor int
	location int
	

	
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
	Println("setter lampe")
	Elev_set_stop_lamp(1)
	
	Println("mekker heis")
	e := new(Elevator)
	e.id = FindElevID()

	//e.stopList = llist.New()
	
	if Elev_get_floor_sensor_signal() == -1 {
		Elev_set_speed(-300)
		for Elev_get_floor_sensor_signal() == -1 {}
	}
	Elev_set_speed(0)
	e.direction = 0
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
	//masterOrderChan := make(chan ButtonSignal, 1)
	//nextDestinationChan := make(chan int, 1)

	go e.OrderHandler(intOrderChan, extOrderChan)
	go Elev_get_order(intOrderChan, extOrderChan)
	go e.Run()
	go e.UpdateStatus()
	go e.printInfo()
	
	Println("ferdig med init")
	return e
}


func (e *Elevator) OrderHandler(intOrderChan chan ButtonSignal, extOrderChan chan ButtonSignal) {
	// external order, send til master
	// internal order legg til først i stopplisten 
	// (kun etter orders som er i etasjer over og i riktig retning)
	Println("orderhandler")
	var newOrder ButtonSignal
	for{
		select{
			case newOrder = <- extOrderChan:
				Println("får ordre")
				Println(newOrder.Floor)
				e.addOrder(newOrder)
			case newOrder = <- intOrderChan:
				e.addOrder(newOrder)
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
				} 
				if e.orderOnCurrentFloor() && e.canCompleteOrder() {
					Elev_set_speed(0)
					if e.currentFloor == 3 || e.currentFloor == 0 {
						e.removeAllOrdersOnFloor(e.currentFloor)
					} else if e.direction > 0 {
						e.removeOrdersGoingUp(e.currentFloor)
					} else if e.direction < 0 {
						e.removeOrdersGoingDown(e.currentFloor)
					} 
					Println("I ORDER ON CURRENT FLOOR: SOVER")
					Sleep(3*Second)
				} else if e.location != -1 {
					if e.orderInCurrentDir() {
						Println("setter speed")
						Elev_set_speed(300*e.direction)
						Sleep(2*Millisecond)
						continue
					} else if e.orderInOtherDir() {
						Println("setter speed og endrer retning")
						Elev_set_speed(300*e.direction)
						e.direction = -e.direction
						Sleep(2*Millisecond)
					} else if e.location != -1 {
						Println("setter direction til null")
						e.direction = 0
					}
				}
				if e.direction == 0 {
					Println("henter ny retning")
					e.getNewDirection()
				}


	}	//for
}	//func


func (e *Elevator) UpdateStatus() {
	for {
		e.location = Elev_get_floor_sensor_signal()
		
		if e.location != -1  && e.currentFloor != e.location{
			e.currentFloor = e.location
			//send reachedfloor
			
		}
		Sleep(Millisecond)
	}
}


func (e *Elevator) addOrder(order ButtonSignal) {
	e.orderMatrix[order.Floor][order.Button - 1] = true
}

func (e *Elevator) removeOrder(floor int, button int) {
	e.orderMatrix[floor][button - 1] = false
}

func (e *Elevator) removeAllOrdersOnFloor(floor int) {
	for i := 1; i <= 3; i++ {
		e.removeOrder(floor, i)
	}
}

func (e *Elevator) removeOrdersGoingUp(floor int) {
	e.removeOrder(floor, BUTTON_CALL_UP)
	e.removeOrder(floor, BUTTON_COMMAND)
}

func (e *Elevator) removeOrdersGoingDown(floor int) {
	e.removeOrder(floor, BUTTON_CALL_DOWN)
	e.removeOrder(floor, BUTTON_COMMAND)
}

func (e *Elevator) getNewDirection() {
		// return int eller send på channel ? - skriv til e.direction
	if (e.orderMatrix[e.currentFloor][0] || e.orderMatrix[e.currentFloor][1] || e.orderMatrix[e.currentFloor][2]) {
		
	}
	dist := N_FLOORS
	next := e.currentFloor
	for i := 0; i <= 3; i++ {
		if e.orderMatrix[i][2] {
			if int(math.Abs(float64(e.currentFloor - i))) < dist {
				dist = int(math.Abs(float64(e.currentFloor - i)))
				next = i
				Println(next)
			}
		}
	}	
	if next == e.currentFloor {
		for i := 0; i <= 3; i++ { // finn nederste som vil opp
			if e.orderMatrix[i][0] {
				dist = int(math.Abs(float64(e.currentFloor - i)))
				next = i
			}	
		}	
		for i := 3; i >= 0; i-- { // finn øverste ordre som vil ned
			if e.orderMatrix[i][1] {
				if int(math.Abs(float64(e.currentFloor - i))) < dist {
					dist = int(math.Abs(float64(e.currentFloor - i)))
					next = i
				}
			}
		}
	}
	if next != e.currentFloor {
		if next > e.currentFloor {
			e.direction = 1
		} else {
			e.direction = -1
		}
		return
	}
	return
}


func (e *Elevator) orderInOtherDir() bool {
	if e.direction > 0 {
		for i := 0; i <= 3; i++ {
			if e.orderMatrix[i][1] {
				return true
			}
		}
	} else if e.direction < 0 {
		for i := 0; i <= 3; i++ {
			if e.orderMatrix[i][0] {
				return true
			}		
		}
	}
	return false
}

func (e *Elevator) orderOnCurrentFloor() bool {	
	if Elev_get_floor_sensor_signal() == -1 {
		return false
	}
	if (e.orderMatrix[e.currentFloor][0] || e.orderMatrix[e.currentFloor][1] || e.orderMatrix[e.currentFloor][2]) {
		Println("returnerer true")
		return true
	}
	return false
}

func (e *Elevator) orderInCurrentDir() bool {
	if e.direction > 0 {
		for i := e.currentFloor; i <= 3; i++ {
			if e.orderMatrix[i][0] || e.orderMatrix[i][2] {
				return true
			}
		}	
	} else if e.direction < 0 {
		for i := e.currentFloor; i >= 0; i-- {
			if e.orderMatrix[i][1] || e.orderMatrix[i][2] {
				return true
			}
		}
	}
	return false
}

func (e *Elevator) canCompleteOrder() bool {
	Println("SJEKKER OM ORDRE KAN GJENNOMFØRES")
	if e.orderMatrix[e.currentFloor][2] || e.currentFloor == 3 || e.currentFloor == 0 {
		return true
	} else if e.orderMatrix[e.currentFloor][0] && e.orderMatrix[e.currentFloor][1] {
		return true
	} else if e.orderMatrix[e.currentFloor][0] && e.direction > 0 {
		return true
	} else if e.orderMatrix[e.currentFloor][1] && e.direction < 0 {
		return true
	} else {
		Println("KUNNE IKKE GJENNOMFØRE ORDRE")
		return false
	}
}

	

func (e *Elevator) printInfo() {
	for {
		Printf("RETNING: %v \n", e.direction)
		Printf("CURRENTFLOOR: %v \n", e.currentFloor)
		Printf("ORDERMATRIX\n\topp    ned    intern\n")
		Printf("3.\t")
		Println(e.orderMatrix[3])
		Printf("2.\t")
		Println(e.orderMatrix[2])
		Printf("1.\t")
		Println(e.orderMatrix[1])
		Printf("0.\t")
		Println(e.orderMatrix[0])
		Sleep(2*Second)
	}
}
