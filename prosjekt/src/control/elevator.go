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
	//locked bool
	speed int
	

	
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

	if Elev_init() != 1 {
		Println("Could not initialize elevator")
	}
	
	e := new(Elevator)
	e.id = FindElevID()

	//e.stopList = llist.New()
	
	if Elev_get_floor_sensor_signal() == -1 {
		e.speed = Elev_set_speed(-300)
		for Elev_get_floor_sensor_signal() == -1 {}
	}
	e.speed = Elev_set_speed(0)
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
	arrivedAtFloorChan := make(chan int, 1)
	getMovingChan := make(chan int, 1)
	//keepMovingChan := make(chan int, 1)

	go e.OrderHandler(intOrderChan, extOrderChan, getMovingChan)
	go Elev_get_order(intOrderChan, extOrderChan)
	go e.Run(arrivedAtFloorChan, getMovingChan/*, keepMovingChan*/)
	go e.UpdateStatus(arrivedAtFloorChan)
	go e.printInfo()
	
	Println("ferdig med init")
	return e
}


func (e *Elevator) OrderHandler(intOrderChan chan ButtonSignal, extOrderChan chan ButtonSignal, getMovingChan chan int /*, fromMasterChan chan ButtonSignal*/) {
	// external order, send til master
	// internal order legg til først i stopplisten 
	// (kun etter orders som er i etasjer over og i riktig retning)
	Println("orderhandler")
	var newOrder ButtonSignal
	for{
		Println("starten av OrderHandler løkke")
		select{
			case newOrder = <- extOrderChan: // slave() leser direkte fra extOrderChan (fjern)
				e.addOrder(newOrder)
				if e.direction == 0 {
					getMovingChan <- 1
				}
				/* SEND MESSAGE TIL MASTER
			case newOrder = <- fromMasterChan:
				e.addOrder(newOrder)
				if e.direction == 0 {
					getMovingChan <- 1
				}
				*/
			case newOrder = <- intOrderChan:
				e.addOrder(newOrder)
				if e.direction == 0 {
					getMovingChan <- 1
				}
		}
	}
}

func (e *Elevator) Run(arrivedAtFloorChan chan int, getMovingChan chan int) {
	// go funksjonen som skriver til nextDestinationChan
	// go funksjonen som oppdaterer elevator variablene
	for{
		Println("STARTEN AV RUN")
		select{
			case <- stopSignalChan: 
				e.speed = Elev_set_speed(0)
				Elev_set_stop_lamp(-1)
				// hva mer skal skje?
			case <- isObstructedChan:
				// handle obstruction
			case <- arrivedAtFloorChan:
						 
				if e.canCompleteOrder() || !e.moreOrdersInCurrentDir() { // del opp if - if else
					e.speed = Elev_set_speed(0)
					// hvis (e.moreOrdersInDir() (sjekker om orderen i etasjen vil samme vei) (bruk orderWithSameDir og inkluder nåværende etasje) -> behold dir
					// hvis (!.moreOrdersInDir())
					Println("I ORDER ON CURRENT FLOOR: SOVER")
					Elev_set_door_open_lamp(1)
					Sleep(2*Second)
					Elev_set_door_open_lamp(0)
					if e.currentFloor == 3 || e.currentFloor == 0 { // lag en egen funksjon for denne delen
						e.removeAllOrdersOnFloor(e.currentFloor) 
					} else if e.direction > 0 {
						e.removeOrdersGoingUp(e.currentFloor) // legg til en keep moving kanal inne i denne for å unngå at det henger i orderHandler
					} else if e.direction < 0 {
						e.removeOrdersGoingDown(e.currentFloor)
					} else {
						if e.currentFloor >= N_FLOORS/2 {
							e.removeOrdersGoingUp(e.currentFloor)
							e.direction = 1
						} else {
							e.removeOrdersGoingDown(e.currentFloor)
							e.direction = -1
						}
					}
				} // dette vil sannsynligvis være ubrukelig
				if e.orderWithSameDir() && e.speed == 0 {
					Println("setter speed")
					e.speed = Elev_set_speed(300*e.direction)
					continue
				} else if e.orderInOtherDir() && e.speed == 0 {
					Println("setter speed og endrer retning")
					e.direction = -e.direction
					if e.canCompleteOrder() {
						arrivedAtFloorChan <- 1
					}
					e.speed = Elev_set_speed(300*e.direction)
				} else if e.speed == 0 {
					e.direction = 0
				} // ubrukelig til hit
			case <- getMovingChan:
				if e.direction == 0 {
					e.getNewDirection(arrivedAtFloorChan)
					e.speed = Elev_set_speed(300*e.direction)							
					/*if e.orderWithSameDir() {
						Println("setter speed")
								
						e.speed = Elev_set_speed(300*e.direction)
								
						continue
					} else if e.orderInOtherDir() {
						Println("setter speed og endrer retning")
						
						e.direction = -e.direction
						e.speed = Elev_set_speed(300*e.direction)
							
					}*//* else if e.speed == 0 {
						e.direction = 0
					}*/
				}

			}
	}	//for
}	//func


func (e *Elevator) UpdateStatus(arrivedAtFloorChan chan int/*, updateChan chan network.ElevatorInfo*/) {
	lastDir = e.direction
	for {
		e.location = Elev_get_floor_sensor_signal()
		
		if e.location != -1  && e.currentFloor != e.location{
			e.currentFloor = e.location
			Elev_set_floor_indicator(e.currentFloor) // fiks så den bare lyser når heisen står stille?
			Printf("----------------------\nI ETASJE %v\n--------------------\n", Elev_get_floor_sensor_signal())
			//send reachedfloor
			arrivedAtFloorChan <- 1
			/* SEND MESSAGE TIL MASTER
			updateChan <- ElevatorInfo{}
			*/
			/* SEND MESSAGE TIL MASTER
		} else if e.direction != lastDir{
			// send update
		}
		*/
		Sleep(50*Millisecond)
		}
	}
}


func (e *Elevator) addOrder(order ButtonSignal) {
	e.orderMatrix[order.Floor][order.Button] = true
	Elev_set_button_lamp(order)
}

func (e *Elevator) removeSingleOrder(floor int, button int) {
	e.orderMatrix[floor][button] = false
	Elev_set_button_lamp(ButtonSignal{button, floor, 0})
}

func (e *Elevator) removeAllOrdersOnFloor(floor int) {
	for i := 0; i <= 2; i++ {
		e.removeSingleOrder(floor, i)
	}
}

func (e *Elevator) removeOrdersGoingUp(floor int) {
	e.removeSingleOrder(floor, BUTTON_CALL_UP)
	e.removeSingleOrder(floor, BUTTON_COMMAND)
}

func (e *Elevator) removeOrdersGoingDown(floor int) {
	e.removeSingleOrder(floor, BUTTON_CALL_DOWN)
	e.removeSingleOrder(floor, BUTTON_COMMAND)
}
/* BRUK I RUN
func (e *Elevator) removeOrders() { 
	if e.currentFloor == 3 || e.currentFloor == 0 { // lag en egen funksjon for denne delen
		e.removeAllOrdersOnFloor(e.currentFloor) 
	} else if e.direction > 0 {
		e.removeOrdersGoingUp(e.currentFloor) // legg til en keep moving kanal inne i denne for å unngå at det henger i orderHandler
	} else if e.direction < 0 {
		e.removeOrdersGoingDown(e.currentFloor)
	} else {
		if e.currentFloor >= N_FLOORS/2 {
			e.removeOrdersGoingUp(e.currentFloor)
			e.direction = 1
		} else {
			e.removeOrdersGoingDown(e.currentFloor)
			e.direction = -1
		}
	}
}
*/
func (e *Elevator) getNewDirection(arrivedAtFloorChan chan int) {
		// return int eller send på channel ? - skriv til e.direction
	if (e.orderMatrix[e.currentFloor][0] || e.orderMatrix[e.currentFloor][1] || e.orderMatrix[e.currentFloor][2]) {
		arrivedAtFloorChan <- 1
		return
	} //FIKS
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
			Printf("SETTER DIRECTION = %v\n", 1)
			e.direction = 1
		} else {
			Printf("SETTER DIRECTION = %v\n", -1)
			e.direction = -1
		}
		return
	}
	return
}


func (e *Elevator) orderInOtherDir() bool {
	if e.direction > 0 {
		for i := 0; i <= 3; i++ {
			if e.orderMatrix[i][1] || e.orderMatrix[i][2] {
				return true
			}
		}
	} else if e.direction < 0 {
		for i := 0; i <= 3; i++ {
			if e.orderMatrix[i][0] || e.orderMatrix[i][2] {
				return true
			}		
		}
	}
	return false
}

func (e *Elevator) moreOrdersInCurrentDir() bool {
	if e.direction > 0 {
		for i := e.currentFloor+1; i <= 3; i++ {
			if e.orderOnFloor(i) {
				return true
			}
		}	
	} else if e.direction < 0 {
		for i := e.currentFloor-1; i >= 0; i-- {
			if e.orderOnFloor(i)  {
				return true
			}
		}
	}
	return false
}



func (e *Elevator) orderOnFloor(floor int) bool {	
	if (e.orderMatrix[floor][0] || e.orderMatrix[floor][1] || e.orderMatrix[floor][2]) {
		return true
	}
	return false
}


func (e *Elevator) orderWithSameDir() bool { // bytte navn til f.eks "moreOrdersInDir"
	if e.direction > 0 {
		for i := e.currentFloor+1; i <= 3; i++ {
			if e.orderMatrix[i][0] || e.orderMatrix[i][2] {
				return true
			} else if i == 3 && e.orderMatrix[i][1] { // bytte dette:
				return true
			} // hit
		} // med dette:
		/*for i := 3; i > e.currentFloor; i-- {
			if e.orderMatrix[i][1] {
				return true
			}
		} */	
	} else if e.direction < 0 {
		for i := e.currentFloor-1; i >= 0; i-- {
			if e.orderMatrix[i][1] || e.orderMatrix[i][2] {
				return true
			} else if i == 0 && e.orderMatrix[i][0] { // bytt dette:
				return true
			} // hit
		} // med dette:
		/*for i := 0; i < e.currentFloor; i++ {
			if e.orderMatrix[i][2] {
				return true
			}
		} */
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
		Printf("LOCATION: %v \n", e.location)
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
