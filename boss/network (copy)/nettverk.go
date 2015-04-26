package nettverk

import (
	."net"
	//"time"
	"encoding/json"
)

const (
	Alive int = iota
	NewOrder
	CompletedOrder
	Cost
)

type Message struct{
	Content int
	Addr string `json:"-"` // <- se nærmere på
	Floor int
	Button int
	Cost int
}



func listenUDP() {
	buf := make([]byte, 1024)
	addr, err := ResolveUDPAddr("udp", ":27346")
	if err != nil {
    	panic(err)
	}
	sock, err := ListenUDP("udp", addr)
	if err != nil {
    	panic(err)
	}
	var msg Message
	for {
		length, senderAddr, err := sock.ReadFromUDP(buf)
		if err != nil {
    		panic(err)
		}	
		json.Unmarshal(buf[:length], &msg)
		var _ = senderAddr
		// send msg videre
	} 
}
	
	
func sendUDP(msg Message) {

	addr, err := ResolveUDPAddr("udp", "129.241.187.255:27346")
	if err != nil {
    	panic(err)
	}
	sock, err := DialUDP("udp", nil, addr)
	if err != nil {
    	panic(err)
	}	

	buf, err := json.Marshal(msg)
	if err != nil {
    	panic(err)
	}

	_, err = sock.Write(buf)
	if err != nil {
    	panic(err)
	}
}



