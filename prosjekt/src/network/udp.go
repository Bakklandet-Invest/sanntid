package network

import (
	"net"
	"strconv"
	"fmt"
)


type messageUDP struct{ 
	recieveAddr string
	data []byte
	length int // length of recieved data in bytes
}

// For testing
func PrintUDPMessage(msg  messageUDP){ 
	fmt.Printf("msg messageUDP: \n recieveAddr = %s \n data = %s \n length = %v \n", msg.recieveAddr, msg.data, msg.length)
}

func InitUDP(localListenPort, broadcastPort, messageSize int, sendCh, recieveCh chan messageUDP)(err error){
	fmt.Println("InitUDP running")
	
	broadcastAddr, err = net.ResolveUDPAddr("udp","255.255.255.255:"+strconv.Itoa(broadcastPort))
	if err != nil{
		return err
	}
	
	// Find localaddress and sets up with choosen port
	tempConn, err := net.DialUDP("udp", nil, broadcastAddr)
	defer tempConn.Close()
	tempAddr := tempConn.LocalAddr()
	LocalAddress, err = net.ResolveUDPAddr("udp", tempAddr.String())
	LocalAddress.Port = localListenPort
	

	localListenConn, err := net.ListenUDP("udp", LocalAddress)
	if err != nil {
		return err
	}
		
	broadcastListenConn, err := net.ListenUDP("udp", broadcastAddr)
	if err != nil {
		localListenConn.Close()
		return err
	}
	
	go recieveServerUDP(localListenConn, broadcastListenConn, messageSize, recieveCh)
	go transmitServerUDP(localListenConn, broadcastListenConn, sendCh)

	return err
}

func recieveServerUDP(lConn, bConn *net.UDPConn, messageSize int, receiveCh chan messageUDP) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in recieveServerUDP: %s \n Closing connection.", r)
			lConn.Close()
			bConn.Close()
		}
	}()

	bConnRecieveCh := make(chan messageUDP)
	lConnRecieveCh := make(chan messageUDP)

	go connectionReaderUDP(lConn, messageSize, lConnRecieveCh)
	go connectionReaderUDP(bConn, messageSize, bConnRecieveCh)

	for {
		select {
		case buffer := <-bConnRecieveCh:
			receiveCh <- buffer

		case buffer := <-lConnRecieveCh:
			receiveCh <- buffer
		}
	}
}

func transmitServerUDP(lConn, bConn *net.UDPConn, sendCh chan messageUDP) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in transmitServerUDP: %s \n Closing connection.", r)
			lConn.Close()
			bConn.Close()
		}
	}()

	var err error
	var n int

	for {
		msg := <-sendCh // Waits for new message
		
		if msg.recieveAddr == "broadcast" {
			n, err = lConn.WriteToUDP(msg.data, broadcastAddr)
		} else {
			recieveAddr, err := net.ResolveUDPAddr("udp", msg.recieveAddr)
			if err != nil {
				fmt.Printf("Error in transmitServerUDP: ResolveUDPAddr failed\n")
				panic(err)
			}
			n, err = lConn.WriteToUDP(msg.data, recieveAddr)
		}
		if err != nil || n < 0 {
			fmt.Printf("Error in transmitServerUDP \n")
			panic(err)
		}
		//fmt.Printf("transmitServerUDP sent %s to %s\n", msg.data, msg.recieveAddr)
	}
}



func connectionReaderUDP(conn *net.UDPConn, messageSize int, rCh chan messageUDP) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error in connectionReaderUDP: %s \n Closing connection.", r)
			conn.Close()
		}
	}()

	buffer := make([]byte, messageSize)

	for {

		n, recieveAddr, err := conn.ReadFromUDP(buffer) // n number of bytes copied to buffer

		if err != nil || n < 0 {
			panic(err)
		}
		rCh <- messageUDP{recieveAddr: recieveAddr.String(), data: buffer, length: n}
	}
}
