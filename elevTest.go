package main

import (
	."driver"
	."fmt"	
	."time"
	)
	
func main() {
	Println(Elev_get_floor_sensor_signal())
	Elev_set_speed(300)
	Sleep(100*Millisecond)
	for Elev_get_floor_sensor_signal() == -1 {}
	Elev_set_speed(0)
	return
}
