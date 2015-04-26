package control

import (
	"math"
	)



func SimpleCost(elevFloor int, orderFloor int) int {
	cost := int(math.Abs(float64(elevFloor - orderFloor)))
	return cost
}

func ComplexCost(elevDirection int, elevFloor int, elevOrderMatrix [4][3]bool, orderFloor int) int { // får bare inn ordere som er i retning, men motsatt rettet, eller "bak" heisen
	cost := 0
	betweenOrders := false
	if elevDirection > 0 {
		if elevFloor < orderFloor { // orderen vil ned
			for i := 3; i > elevFloor; i-- {
				if !orderOnFloor(i, elevOrderMatrix) && !betweenOrders {
					continue
				} else {
					if elevOrderMatrix[i][0] {
						cost++
					}
					if elevOrderMatrix[i][1] && i < orderFloor { // tar bare med orderene ned som heisen stopper på før orderen
						cost++
					}
					if elevOrderMatrix[i][2] && !(elevOrderMatrix[i][0] || elevOrderMatrix[i][1]) {
						cost++
					}
					betweenOrders = true
				}	
				cost++
			}

		} else { // orderen ligger "bak" heisen ---- skal ikke legge til andre ordre i samme retning som er mellom heisen og ordre, eller etasjer den ikke skal passere (alt annet)
			for i := 0; i <= 3; i++ {
				if !orderOnFloor(i, elevOrderMatrix) && (!betweenOrders || !isOrdersAbove(i, elevOrderMatrix)) {
					continue
				} else {
					if elevOrderMatrix[i][0] && (i >= elevFloor && i < orderFloor) { // tar bare med orderene ned som heisen stopper på før orderen
						cost++
					}
					if elevOrderMatrix[i][1] {
						cost++
					}
					if elevOrderMatrix[i][2] && !(elevOrderMatrix[i][0] || elevOrderMatrix[i][1]) {
						cost++
					}
					betweenOrders = true
				}	
				cost++
			}
		}
	} else { //elev.direction < 0
		if elevFloor > orderFloor { // orderen vil opp
			for i := 3; i < elevFloor; i-- {
				if !orderOnFloor(i, elevOrderMatrix) && !betweenOrders {
					continue
				} else {
					if elevOrderMatrix[i][0] && i < orderFloor { // tar bare med orderene opp som heisen stopper på før orderen
						cost++
					}
					if elevOrderMatrix[i][1] {
						cost++
					}
					if elevOrderMatrix[i][2] && !(elevOrderMatrix[i][0] || elevOrderMatrix[i][1]) {
						cost++
					}
				}
				cost++
			}
		} else { // orderen ligger "bak" heisen ---- skal ikke legge til andre ordre i samme retning som er mellom heisen og ordre, eller etasjer den ikke skal passere (alt annet)
			for i := 0; i <= 3; i++ {
				if !orderOnFloor(i, elevOrderMatrix) && (!betweenOrders || !isOrdersBelow(i, elevOrderMatrix)) {
					continue
				} else {
					if elevOrderMatrix[i][0] {
						cost++
					}
					if elevOrderMatrix[i][1] && (i >= elevFloor && i < orderFloor) { // tar bare med orderene ned som heisen stopper på før orderen
						cost++
					}
					if elevOrderMatrix[i][2] && !(elevOrderMatrix[i][0] || elevOrderMatrix[i][1]) {
						cost++
					}
					betweenOrders = true
				}	
				cost++
			}
		}
	}
	return cost
}

func orderOnFloor(floor int, orderMatrix [4][3]bool) bool {
	return (orderMatrix[floor][0] || orderMatrix[floor][1] || orderMatrix[floor][2])
}

func isOrdersAbove(floor int, orderMatrix [4][3]bool) bool {
	for i := 3; i > floor; i-- {
		if orderOnFloor(floor, orderMatrix) {
			return true
		}
	}
	return false
}

func isOrdersBelow(floor int, orderMatrix [4][3]bool) bool {
	for i := 0; i < floor; i++ {
		if orderOnFloor(floor, orderMatrix) {
			return true
		}
	}
	return false
}
