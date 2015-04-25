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
	speed int

	//leaveCheck int
	

	
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

func InitElevator(updateOutChan chan network.ElevatorInfo, updateInChan chan network.ElevatorInfo, checkMasterChan chan string, newOrderChan chan ButtonSignal, completeOrderChan chan ButtonSignal, fromMasterChan chan ButtonSignal, intOrderChan chan ButtonSignal, extOrderChan chan ButtonSignal ) *Elevator {

	if Elev_init() != 1 {
		Println("Could not initialize elevator")
	}
	
	e := new(Elevator)
	e.id = FindElevID()

	
	if Elev_get_floor_sensor_signal() == -1 {
		e.speed = Elev_set_speed(-300)
		for Elev_get_floor_sensor_signal() == -1 {}
	}
	e.speed = Elev_set_speed(0)
	e.direction = 0
	e.currentFloor = Elev_get_floor_sensor_signal()


	/* CHANALS SOM ER INPUT */
	/*
	updateOutChan := make(chan network.ElevatorInfo)
	updateInChan := make(chan network.ElevatorInfo)
	checkMasterChan := make(chan string)
	newOrderChan := make(chan ButtonSignal)
	completeOrderChan := make(chan ButtonSignal)
	fromMasterChan := make(chan ButtonSignal)
	*/

	//intOrderChan := make(chan ButtonSignal, 1) // brukes i solo
	//extOrderChan := make(chan ButtonSignal, 1) // brukes i solo
	

	/* CHANALS SOM ER INTERNE */
	arrivedAtFloorChan := make(chan int, 1) 
	getMovingChan := make(chan int, 1)

	go e.OrderHandler(intOrderChan, extOrderChan, fromMasterChan)
	go Elev_get_order(intOrderChan, extOrderChan)
	go e.Run(arrivedAtFloorChan, getMovingChan, completeOrderChan)
	go e.UpdateStatus(arrivedAtFloorChan, updateOutChan)
	go e.printInfo()
	
	Println("ferdig med init")
	return e
}


func (e *Elevator) OrderHandler(intOrderChan chan ButtonSignal, fromMasterChan chan ButtonSignal, getMovingChan chan int) {
	Println("orderhandler")
	var newOrder ButtonSignal
	for{
		Println("starten av OrderHandler løkke")
		select{
			case newOrder = <- fromMasterChan:
				e.addOrder(newOrder)
				if e.direction == 0 {
					getMovingChan <- 1
				}
			case newOrder = <- intOrderChan:
				e.addOrder(newOrder)
				if e.direction == 0 {
					getMovingChan <- 1
				}
		}
	}
}

func (e *Elevator) Run(arrivedAtFloorChan chan int, getMovingChan chan int, completeOrderChan chan ButtonSignal) {
	// go funksjonen som skriver til nextDestinationChan
	// go funksjonen som oppdaterer elevator variablene
	for{
		Println("STARTEN AV RUN")
		select{
			/*case <- stopSignalChan:
				//e.speed = Elev_set_speed(0)
				e.speed = e.stopElevator()
				Elev_set_stop_lamp(-1)
				// hva skal skje når stop-knappen trykkes?
			 */
			/*case <- obstuctionChan:
				// handle obstruction 
			*/
			case <- arrivedAtFloorChan:
				if e.canCompleteOrder()  || !e.moreOrdersInDir() {
					//e.speed = Elev_set_speed(0)
					e.speed = e.stopElevator()
					
					if e.moreOrdersInDir() {
						// e.direction = e.direction
					} else if e.orderInOtherDir() {
						e.direction = -e.direction
					} else {
						e.direction = 0
					}
					
					Println("I ORDER ON CURRENT FLOOR: SOVER")
					Elev_set_door_open_lamp(1)
					Sleep(2*Second)
					Elev_set_door_open_lamp(0)

					e.removeOrders(completeOrderChan)

					if e.moreOrdersInDir() {
							// e.direction = e.direction
					} else if e.orderInOtherDir() {
						e.direction = -e.direction
					} else {
						e.direction = 0
					}

					if e.orderOnCurrentFloorInDir() {
						arrivedAtFloorChan <- 1
						continue						
					} else {
						e.speed = Elev_set_speed(300*e.direction)
					}
			}
			
			
			case <- getMovingChan:
				if e.direction == 0 {
					e.getNewDirection(arrivedAtFloorChan)
					e.speed = Elev_set_speed(300*e.direction)							
				}

			}
	}	//for
}	//func


func (e *Elevator) UpdateStatus(arrivedAtFloorChan chan int, updateOutChan chan network.ElevatorInfo) {
	for {
		e.location = Elev_get_floor_sensor_signal()
		
		if e.location != -1  && e.currentFloor != e.location{
			e.currentFloor = e.location
			Elev_set_floor_indicator(e.currentFloor) // fiks så den bare lyser når heisen står stille?
			Printf("----------------------\nI ETASJE %v\n--------------------\n", Elev_get_floor_sensor_signal())
			//send reachedfloor
			arrivedAtFloorChan <- 1
			updateOutChan <- network.ElevatorInfo{e.orderMatrix, e.currentFloor, e.direction}
			
		}/* else if e.location == -1 && e.location != e.leaveCheck {
			// sjekk for å sende update når heisen forlater en etasje
		}*/
		Sleep(50*Millisecond)
	}
}

func (e *Elevator) stopElevator () int {
	if e.speed == 0 {
		return 0
	} else {
		Elev_set_speed(-300*e.direction)
		Sleep(5*Millisecond)
	} 
	Elev_set_speed(0)
	return 0
}

func (e *Elevator) addOrder(order ButtonSignal) {
	e.orderMatrix[order.Floor][order.Button] = true
	Elev_set_button_lamp(order)
}

func (e *Elevator) removeSingleExtOrder(floor int, button int, completeOrderChan chan ButtonSignal) {
	e.orderMatrix[floor][button] = false
	Elev_set_button_lamp(ButtonSignal{button, floor, 0}) // lag en lokal "completedOrder" variabel?
	completedOrderChan <- {ButtonSignal{button, floor, 0}
}

func (e *Elevator) removeSingleIntOrder(floor int, button int) { //mulig denne er ubrukelig hvis master styrer lys
	e.orderMatrix[floor][button] = false
	Elev_set_button_lamp(ButtonSignal{button, floor, 0})
}

func (e *Elevator) removeAllOrdersOnFloor(floor int, completeOrderChan chan ButtonSignal) {
		e.removeSingleExtOrder(floor, BUTTON_CALL_UP, completeOrderChan)
		e.removeSingleExtOrder(floor, BUTTON_CALL_DOWN, completeOrderChan)
		e.removeSingleIntOrder(floor, BUTTON_COMMAND)
}

func (e *Elevator) removeOrdersGoingUp(floor int, completeOrderChan chan ButtonSignal) {
	e.removeSingleExtOrder(floor, BUTTON_CALL_UP, completeOrderChan)
	e.removeSingleIntOrder(floor, BUTTON_COMMAND)
}

func (e *Elevator) removeOrdersGoingDown(floor int, completeOrderChan chan ButtonSignal) {
	e.removeSingleExtOrder(floor, BUTTON_CALL_DOWN, completeOrderChan)
	e.removeSingleIntOrder(floor, BUTTON_COMMAND)
}

func (e *Elevator) removeOrders(completeOrderChan chan ButtonSignal) { 
	if e.currentFloor == 3 || e.currentFloor == 0 { // lag en egen funksjon for denne delen
		e.removeAllOrdersOnFloor(e.currentFloor, completeOrderChan) 
	} else if e.direction > 0 {
		e.removeOrdersGoingUp(e.currentFloor, completeOrderChan) // legg til en keep moving kanal inne i denne for å unngå at det henger i orderHandler
	} else if e.direction < 0 {
		e.removeOrdersGoingDown(e.currentFloor, completeOrderChan)
	} else {
		if e.currentFloor >= N_FLOORS/2 {
			e.removeOrdersGoingUp(e.currentFloor, completeOrderChan)
		} else {
			e.removeOrdersGoingDown(e.currentFloor, completeOrderChan)
		}
	}
}

func (e *Elevator) getNewDirection(arrivedAtFloorChan chan int) {
		// return int eller send på channel ? - skriv til e.direction
	if e.orderOnFloor(e.currentFloor) {
		if e.orderMatrix[e.currentFloor][0] && e.orderMatrix[e.currentFloor][1] {
			if e.currentFloor >= N_FLOORS/2 {
				Printf("SETTER DIRECTION = %v\n", 1)
				e.direction = 1
			} else {
				Printf("SETTER DIRECTION = %v\n", -1)
				e.direction = -1
			}
		} else if e.orderMatrix[e.currentFloor][0] {
			Printf("SETTER DIRECTION = %v\n", 1)
			e.direction = 1
		} else if e.orderMatrix[e.currentFloor][1] {
			Printf("SETTER DIRECTION = %v\n", -1)
			e.direction = -1
		} else {
			e.removeSingleOrder(e.currentFloor, BUTTON_COMMAND)			
			return
		}
		arrivedAtFloorChan <- 1
		return
	}							 //FIKS
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


func (e *Elevator) orderOnCurrentFloorInDir() bool {
	if e.direction > 0 {
		return e.orderMatrix[e.currentFloor][0]
	} else if e.direction < 0 {
		return e.orderMatrix[e.currentFloor][1]
	} 
	return false
}


func (e *Elevator) orderInOtherDir() bool { // returnerer true hvis heisen skal bytte dir
	if e.direction > 0 {
		if e.orderMatrix[e.currentFloor][1] {
			return true
		} else {
			for i := 0; i < e.currentFloor; i++ {
				if e.orderOnFloor(i) {
					return true
				}
			}
		}		
	} else if e.direction < 0 {
		if e.orderMatrix[e.currentFloor][0] {
			return true
		} else {
			for i := 3; i > e.currentFloor; i-- {
				if e.orderOnFloor(i) {
					return true
				}		
			}
		}
	} 
	return false
}

func (e *Elevator) orderInCurrentDir() bool {
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


func (e *Elevator) moreOrdersInDir() bool {
	if e.direction > 0 {
		if e.currentFloor == 3 {
			return false
		} else if e.orderMatrix[e.currentFloor][0] {
			return true
		} else {
			for i := e.currentFloor+1; i <= 3; i++ {
				if e.orderOnFloor(i) {
					return true
				}
			}
		}		
	} else if e.direction < 0 {
		if e.currentFloor == 0 {
			return false
		} else if e.orderMatrix[e.currentFloor][1] {
			return true
		} else {
			for i := 0; i < e.currentFloor; i++ {
				if e.orderOnFloor(i) {
					return true
				}
			}
		}
	}
	return false
}



func (e *Elevator) canCompleteOrder() bool {
	Println("SJEKKER OM ORDRE KAN GJENNOMFØRES")
	if e.direction == 0 {
		return true	
	} else if e.orderMatrix[e.currentFloor][2] || e.currentFloor == 3 || e.currentFloor == 0 {
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

