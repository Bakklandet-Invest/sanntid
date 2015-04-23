package main

import(
	"network"
	//"net"
	"fmt"
	"time"
	)
	
func main(){
	var kanal = make(chan int)
	
	network.Init()
	testmessage := network.Message{Content: network.NewOrder, Floor: 3, Button: 3, Cost: 69}
	for {
		fmt.Println(localListenPort)
		network.MessageCh <- testmessage
		time.Sleep(5*time.Second)
	}
	
	<-kanal
}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	//var kanal = make(chan int)
	//var broadcastAddr *net.UDPAddr 
	
	//broadcastAddr, _ = net.ResolveUDPAddr("udp","255.255.255.255:13337")
	//tempConn, err := net.DialUDP("udp", nil, broadcastAddr)
	//if err != nil{
		//panic(err)
	//}
	//defer tempConn.Close()
	//tempAddr := tempConn.LocalAddr()
	//LocalAddress, err := net.ResolveUDPAddr("udp", tempAddr.String())
	//if err != nil{
		//panic(err)
	//}
	//fmt.Println("Localaddress:",LocalAddress)
	
	//LocalAddress.Port = 16969
	
	//fmt.Println("Localaddress med ny port:",LocalAddress)
	
	
	
	//network.Init()
	//<-kanal
//}
	
