package control

import (
	//"llist"
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
	internalOrderChan chan int
	externalOrderChan chan int
	nextDestinationChan chan int
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
	//e.stopList = llist.New()
	e.update = make(chan Elevator, 1)
	e.order = make(chan int, 1)
	
	if Elev_get_floor_sensor_signal() == -1 {
		Elev_set_speed(-300)
		for Elev_get_floor_sensor_signal() == -1 {}
	}
	Elev_set_speed(0)
	e.currentFloor = Elev_get_floor_sensor_signal()
	return e
}




func (elev *Elevator) Run() {
	// go funksjonen som skriver til nextDestinationChan
	// go funksjonen som oppdaterer elevator variablene
	for{
		/* if Elev_get_obstruction_signal():
				handle obstruction */
		//else:
		select{
			case elev.destination = <-nextDestinationChan: // lag funksjon som skriver neste
				if elev.destination == currentFloor {
					e.direction =
				}


		}	//select	
	}	//for
}	//func