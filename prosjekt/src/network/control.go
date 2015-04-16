package control



type Control struct {
	elevators map[string]*Elevator
	orders map[int][2]bool
	connected map[string]bool

}

func (cont *Control) InitMaps() {
	
	c.elevators = make(map[string]*Elevator)
	c.connected = make(map[string]bool)

	c.orders = make(map[int][2]bool)
	
	return
}


master lager en kontrollstruct som lagrer heisID som er siste 3 siffer i IPen, 
IPer, IPen til masterheisen, posisjon, retning osv. Alle heiser lagrer denne. 