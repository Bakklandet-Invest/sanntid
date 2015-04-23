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

	if e.stopList.Front() != nil { // hvis listen er tom
		e.stoplist.PushFront(incomingOrder) 
	} else {
		nextElem := e.stopList.Front()
		root := nextElem.Prev()
		
			if (incomingOrder.Floor == Elev_get_floor_sensor_signal() && e.orderInCurrentDir(incomingOrder.button)){ // ordre i etasjen heisen står
				//få heisen til å ta orderen med en gang
			}
			else if (incomingOrder.button == BUTTON_CALL_COMMAND){ // intern ordre - har lagt til funksjonen til en ekstern ordre med bestemt retning. fiks
				if(e.direction > 0){ // heisen er på vei opp
					if(incomingOrder.Floor > e.currentFloor) { // etasjen ligger over nåværende
						for incomingOrder.Floor > nextElem.Value.(ButtonSignal).Floor { // går så lenge incOrder er større enn neste i køen
							if (nextElem.Value.(ButtonSignal).Floor > nextElem.Next().Value.(ButtonSignal).Floor) { // hvis heisen skal snu, legg til før den snur
								e.stopList.InsertBefore(incomingOrder, nextElem.Next())
								continue // gå ut av forløkken
							}
							nextElem = nextElem.Next() // gå til neste element i forløkken
						}
						e.stopList.InsertBefore(incomingOrder, nextElem.Next()) // hvis incOrder < nextElem
					}
					else {
						// hvis inter ordre ligger under heisen når den er på vei opp
						nextElem = root.Prev() // setter nextElem til det siste elmentet
						if (nextElem.Value.(ButtonSignal).Floor > nextElem.Prev().Value.(ButtonSignal).Floor) { // hvis heisen skal opp igjen
							for (nextElem.Value.(ButtonSignal).Floor > nextElem.Prev().Value.(ButtonSignal).Floor) { // så lenge neste er lavere i listen
								if (incomingOrder.Floor < nextElem.Value.(ButtonSignal).Floor && incomingOrder.Floor > nextElem.Prev().Value.(ButtonSignal).Floor) {
								// Hvis incOrder ligger mellom de to neste orderene i listen
								e.stopList.InsertBefore(incomingOrder, nextElem)
								}
								nextElem.Prev() // fortsett bakover i listen
							}
						}
						e.stopList.InsertBefore(incomingOrder, nextElem)
					}
				}
				else{ // heisen er på vei ned
					if(incomingOrder.Floor < e.currentFloor) { // etasjen ligger under nåværende
						for incomingOrder.Floor < nextElem.Value.(ButtonSignal).Floor { // går så lenge incOrder er mindre enn neste i køen
							if (nextElem.Value.(ButtonSignal).Floor < nextElem.Next().Value.(ButtonSignal).Floor) { // hvis heisen skal snu, legg til før den snur
								e.stopList.InsertBefore(incomingOrder, nextElem.Next())
								continue // gå ut av forløkken
							}
							nextElem = nextElem.Next() // gå til neste element i forløkken
						}
						e.stopList.InsertBefore(incomingOrder, nextElem) // hvis incOrder < nextElem
					}
					else {
						// hvis intern ordre ligger over heisen når den er på vei ned
						nextElem = root.Prev() // setter nextElem til det siste elmentet
						if (nextElem.Value.(ButtonSignal).Floor < nextElem.Prev().Value.(ButtonSignal).Floor) { // hvis heisen skal ned igjen
							for (nextElem.Value.(ButtonSignal).Floor > nextElem.Prev().Value.(ButtonSignal).Floor) { // så lenge neste i listen er høyere
								if (incomingOrder.Floor > nextElem.Value.(ButtonSignal).Floor && incomingOrder.Floor < nextElem.Prev().Value.(ButtonSignal).Floor) {
								// Hvis incOrder ligger mellom de to neste orderene i listen
								e.stopList.InsertBefore(incomingOrder, nextElem)
								}
								nextElem.Prev() // fortsett bakover i listen
							}
						}
						e.stopList.InsertBefore(incomingOrder, nextElem)
					}	
				}
			}
			else if (incomingOrder.button == BUTTON_CALL_UP) { // external order opp
				
				if (e.direction > 0) {
					//for nextElem != root { // unødvendig?
					if (incomingOrder.Floor > e.currentFloor) { // incOrder ligger over heisen
						if (incomingOrder.Floor < nextElem.Value.(ButtonSignal).Floor){ //incOrder ligger mellom current og neste i listen
							e.stopList.InsertBefore(incomingOrder, nextElem)
						}
						else // (incomingOrder.Floor > nextElem.Value.(ButtonSignal).Floor) { //incOrder ligger over current og neste i listen
							{
							//nextElem = nextElem.Next
							for incomingOrder.Floor > nextElem.Value.(ButtonSignal).Floor { // så lenge incOrder er høyere enn neste element
								if (nextElem.Value.(ButtonSignal).Floor > nextElem.Next().Value.(ButtonSignal).Floor) {
									// når neste destinasjon ligger på vei nedover igjen
									e.stopList.InsertBefore(incomingOrder, nextElem.Next())
									continue
								}
								nextElem = nextElem.Next()		
							}
						// trengs det å legge til her? vil aldri være lik, hvis alt annet fungerer
						}
						//nextElem = nextElem.Next()
					}
					else { //incOrder ligger under heisen
						nextElem = root.Prev() // setter nextElem til det siste elmentet
						if (nextElem.Value.(ButtonSignal).Floor > nextElem.Prev().Value.(ButtonSignal).Floor) { // hvis heisen skal opp igjen
							for (nextElem.Value.(ButtonSignal).Floor > nextElem.Prev().Value.(ButtonSignal).Floor) { // så lenge neste er lavere i listen
								if (incomingOrder.Floor < nextElem.Value.(ButtonSignal).Floor && incomingOrder.Floor > nextElem.Prev().Value.(ButtonSignal).Floor)
								// Hvis incOrder ligger mellom de to neste orderene i listen
								e.stopList.InsertBefore(incomingOrder, nextElem)
								nextElem.Prev() // fortsett bakover i listen
							}
						}
						e.stopList.InsertBefore(incomingOrder, nextElem) // hvis heisen ikke har noen kø oppover igjen etter å ha byttet retning
					}
					// } // kommentert for
				}
				else { // e.direction < 0

				}
			}
			else if (incomingOrder.button == BUTTON_CALL_UP) { // incOrder skal opp
				if (e.direction > 0){

				}
			}

			
 // fiks sluttbrackets
	}

}
/* list funksjoner:
func (l *LinkedList) InsertBefore(v interface{}, mark *Element) *Element {
func (l *LinkedList) PushFront(v interface{}) *Element {
func (l *LinkedList) Remove(e *Element) interface{} {
func (l *LinkedList) PushBack(v interface{}) *Element {
func (e *Element) Next() *Element {

Button struct:
	Button int
	Floor int
	Light int
*/


func (e *Elevator) orderInCurrentDir(buttonType int) bool {
	if ((buttonType == BUTTON_CALL_UP && e.direction < 0) || (buttonType = BUTTON_CALL_DOWN && e.direction > 0)) {return false}
	else {return true}
}