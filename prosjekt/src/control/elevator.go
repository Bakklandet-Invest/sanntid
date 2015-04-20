package control

import (
	"llist"
	"strconv"
	"net"
	."driver"
	."fmt"
	)


type Elevator struct {
	id int
	//stopList llist.LinkedList
	direction int
	currentFloor int
	destination int

	updateChan chan Elevator
	intOrderChan chan int
	extOrderChan chan int
	nextDestinationChan chan int
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
	e.id = findElevID()
	e.stopList = llist.New()
	e.update = make(chan Elevator, 1)
	e.order = make(chan int, 1)
	
	if Elev_get_floor_sensor_signal() == -1 {
		Elev_set_speed(-300)
		for Elev_get_floor_sensor_signal() == -1 {}
	}
	Elev_set_speed(0)
	e.currentFloor = Elev_get_floor_sensor_signal()
-

	e.updateChan = make(chan Elevator)
	e.intOrderChan = make(chan ButtonSignal)
	e.extOrderChan = make(chan ButtonSignal)
	e.nextDestinationChan = make(chan int)

	go e.OrderHandler
	go Elev_get_order(intOrderChan, extOrderChan)

	return e
}


func (e *Elevator) OrderHandler() {
	// external order, send til master
	// internal order legg til først i stopplisten 
	// (kun etter orders som er i etasjer over og i riktig retning)
	var newOrder ButtonSignal
	for{
		select{
			case newOrder = <- e.extOrderChan:
				// ..
			case newOrder = <- e.intOrderChan:
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
					ElevSetStopLamp(-1)
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
func (e *Elevator) GetNextDestination() int {
	// skal returnere neste ordre/destinasjon heisen skal til
	if e.stopList.Front() != nil {
		return e.stopList.Front().Value.(int)
	} else {
		return -1
	}
}

func (e *Elevator) addOrder() {
	var incomingOrder ButtonSignal
	
	// få inn ordre fra melding eller fra intOrder og lagre den i incomingOrder

	/* hvis listen er tom, legg til først.
	 hvis det ligger det noe i listen:
	 		- sjekk hvilken retning 
	 			(knapptype/nåværende etasje vs ordre etasje)
			- sjekk etasje
	*/
	if e.stopList.Front() != nil
		e.stoplist.PushFront(incomingOrder)

}
/* list funksjoner:
func (l *LinkedList) InsertBefore(v interface{}, mark *Element) *Element {
func (l *LinkedList) PushFront(v interface{}) *Element {
func (l *LinkedList) Remove(e *Element) interface{} {
func (l *LinkedList) PushBack(v interface{}) *Element {

Button struct:
	Button int
	Floor int
	Light int
*/