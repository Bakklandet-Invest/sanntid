package network

import (
	"net"
	"strings"
)

const NumButtons = 3
const NumFloors = 4

const (
	Alive int = iota + 1
	NewOrder
	CompletedOrder
	Cost
)

// Only message sent over the network
type Message struct{
	Content int
	Addr string
	Floor int
	Button int
	Cost int
}

type ConnectionUDP struct{
	Addr string
	Timer *time.Timer
}


var LocalAddress *net.UDPAddr

var MessageCh = make(chan Message) 
var SyncLightChan = make(chan bool)

// Finds the last quadrant of the IP address
func FindID(ip string) string {
	return strings.Split(strings.Split(ip, ".")[3], ":")[0]
}
