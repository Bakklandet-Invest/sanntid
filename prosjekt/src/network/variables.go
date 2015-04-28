package network

import (
	"net"
	"strings"
	"time"
)

const NumButtons = 3
const NumFloors = 4

const (
	Alive int = iota + 1
	NewOrder
	CompletedOrder
	Info
	Light
)

type ElevatorInfo struct{
	Matrix [NumFloors][NumButtons]bool
	Floor int
	Dir int
}
	

// Only message sent over the network
type Message struct{
	Content int
	Addr string
	Floor int //Floor hvor knappen er trykket
	Button int
	//Cost int
	ElevInfo ElevatorInfo
}

type ConnectionUDP struct{
	Addr string
	Timer *time.Timer
}

const ResetConnTime = 9*time.Second
const NotifyInterval = 6*time.Second

var broadcastAddr *net.UDPAddr

var LocalAddress *net.UDPAddr

var RecieveChan = make(chan messageUDP) //Public
var sendChan = make(chan messageUDP)

var MessageChan = make(chan Message) 
var SyncLightChan = make(chan bool)

// Finds the last quadrant of the IP address
func FindID(ip string) string {
	return strings.Split(strings.Split(ip, ".")[3], ":")[0]
}
