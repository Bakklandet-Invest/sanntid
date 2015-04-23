package main

import (
	."driver"
	."fmt"	
	."time"
	)
	
func main() {
	Elev_init()

	Println(Elev_get_floor_sensor_signal())
	
	Elev_set_speed(-300)
	Println("speed set")	
	Sleep(Second)
	Println("ferdig med sleep")
	for Elev_get_floor_sensor_signal() == -1 {
	}
	Println(Elev_get_floor_sensor_signal())	
	Elev_set_speed(300)
	Sleep(5*Millisecond)
	Elev_set_speed(0)
	Println("heis stoppet")	
	Sleep(Second)
	Println(Elev_get_floor_sensor_signal())
	return
}
