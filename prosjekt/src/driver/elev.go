package driver

import (
	"math"
	."time"
	//."fmt"
)



const (
	BUTTON_CALL_UP int = 0
	BUTTON_CALL_DOWN int = 1
	BUTTON_COMMAND int = 2

	N_BUTTONS int = 3
	N_FLOORS int = 4
	)


type ButtonSignal struct {
	Button int
	Floor int
	Light int
}

var lamp_channel_matrix = [N_FLOORS][N_BUTTONS]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var button_channel_matrix = [N_FLOORS][N_BUTTONS]int{
	{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
	{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
	{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
	{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

func Elev_init() int {
	if !Io_init() {
		return 0
	}

	for i := 0; i < N_FLOORS; i++ {
		if i != 0 {
			Elev_set_button_lamp(ButtonSignal{BUTTON_CALL_DOWN, i, 0})
		}
		if i != N_FLOORS-1 {
			Elev_set_button_lamp(ButtonSignal{BUTTON_CALL_UP, i, 0})
		}
		Elev_set_button_lamp(ButtonSignal{BUTTON_COMMAND, i, 0})
	}

	Elev_set_stop_lamp(0)
	Elev_set_door_open_lamp(0)
	Elev_set_floor_indicator(0)

	return 1
}

func Elev_set_speed(speed int) int {
	last_speed := 0
	if speed > 0 {
		Io_clear_bit(MOTORDIR)
	} else if speed < 0 {
		Io_set_bit(MOTORDIR)
	} else if last_speed < 0 {
		Io_clear_bit(MOTORDIR)
	} else if last_speed > 0 {
		Io_set_bit(MOTORDIR)
	}
	last_speed = speed
 	Io_write_analog(MOTOR, int(2048+4*math.Abs(float64(speed))))
	return speed
}


func Elev_get_floor_sensor_signal() int {
	if Io_read_bit(SENSOR1) {
		return 0
	} else if Io_read_bit(SENSOR2) {
		return 1
	} else if Io_read_bit(SENSOR3) {
		return 2
	} else if Io_read_bit(SENSOR4) {
		return 3
	} else {
		return -1
	}
}

func Elev_get_button_signal(button int, floor int) int {
	// Need error handling before proceeding
	if Io_read_bit(button_channel_matrix[floor][int(button)]) {
		return 1
	} else {
		return 0
	}
}
/*
func (sig *ButtonSignal) ClearPrevButtonSig() {
	time.Sleep(time.Second)
	sig.Floor = -2
}
*/

func Elev_get_order(intOrderChan chan ButtonSignal, extOrderChan chan ButtonSignal) {

	var buttonSig ButtonSignal

	//var prevButtonSig ButtonSignal 
	//prevButtonSig.Floor = -2  // for å unngå at samme ordere sendes mange ganger på kort tid

	for{
		for i := 0; i < 3; i++ {
			if (Elev_get_button_signal(BUTTON_CALL_UP, i) == 1) {
				buttonSig.Floor =  i
				buttonSig.Button = BUTTON_CALL_UP
				buttonSig.Light = 1
				extOrderChan <- buttonSig
			} else if (Elev_get_button_signal(BUTTON_CALL_DOWN, i+1) == 1) {
				buttonSig.Floor =  i+1
				buttonSig.Button = BUTTON_CALL_DOWN
				buttonSig.Light = 1
				extOrderChan <- buttonSig
			} 
		}

		for i := 0; i < 4; i++ {
        
			if ( Elev_get_button_signal( BUTTON_COMMAND, i ) == 1 ) {
				buttonSig.Floor =  i
				buttonSig.Button = BUTTON_COMMAND
				intOrderChan <- buttonSig
			}
		}
		Sleep(30*Millisecond)
	//go prevButtonSig.ClearPrevButtonSig()
	}
}


func Elev_get_stop_signal() bool {
	return Io_read_bit(STOP)
}

func Elev_get_obstruction_signal() bool {
	return Io_read_bit(OBSTRUCTION)
}

func Elev_set_floor_indicator(floor int) {
	// Need error handling before proceeding
	switch floor {
	case 0:
		Io_clear_bit(LIGHT_FLOOR_IND1)
		Io_clear_bit(LIGHT_FLOOR_IND2)
	case 1:
		Io_clear_bit(LIGHT_FLOOR_IND1)
		Io_set_bit(LIGHT_FLOOR_IND2)
	case 2:
		Io_set_bit(LIGHT_FLOOR_IND1)
		Io_clear_bit(LIGHT_FLOOR_IND2)
	case 3:
		Io_set_bit(LIGHT_FLOOR_IND1)
		Io_set_bit(LIGHT_FLOOR_IND2)
	}
}

func Elev_set_button_lamp(order ButtonSignal) {
	// Need error handling before proceeding
	if order.Light == 1 {
		Io_set_bit(lamp_channel_matrix[order.Floor][int(order.Button)])
	} else {
		Io_clear_bit(lamp_channel_matrix[order.Floor][int(order.Button)])
	}
}

func Elev_set_stop_lamp(value int) {
	if value == 1 {
		Io_set_bit(LIGHT_STOP)
	} else {
		Io_clear_bit(LIGHT_STOP)
	}
}

func Elev_set_door_open_lamp(value int) {
	if value == 1 {
		Io_set_bit(LIGHT_DOOR_OPEN)
	} else {
		Io_clear_bit(LIGHT_DOOR_OPEN)
	}
}

func Elev_stop_elev(dir int) int {
	if dir > 0 {
		Elev_set_speed(-300)
		Sleep(5*Microsecond)
	} else if dir < 0 {
		Elev_set_speed(300)
		Sleep(5*Microsecond)
	} 
	Elev_set_speed(0)
	return 0
}
