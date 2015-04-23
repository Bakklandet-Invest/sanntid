package main

import(
	"network"
	"time"
	"fmt"
)

var liftsOnline = make(map[string]network.ConnectionUDP)

func writemap() {
	for {
		fmt.Println("liftsOnline:",liftsOnline)
		time.Sleep(time.Second*10)
	}
}

func main(){
	network.Init()
	go writemap()

	for{
		select{
		case msg := <-network.RecieveChan:
			messageHandler(network.ParseMessage(msg))
		}
			
			
	}
}

func messageHandler(msg network.Message) {
	switch msg.Content{
		case network.Alive:
			if conn, alive := liftsOnline[msg.Addr]; alive{
				conn.Timer.Reset(network.ResetConnTime)
			} else{
				newConn := network.ConnectionUDP{msg.Addr, time.NewTimer(network.ResetConnTime)}
				liftsOnline[msg.Addr] = newConn
			}
		case network.NewOrder:
			//kanskje skrive ut noe?
			
			cost :=  69 //costfunksjonen inn her
			
			costMsg := network.Message{Content: network.Cost, Floor: msg.Floor, Button: msg.Button, Cost: cost}
			network.MessageCh <- costMsg
		case network.CompletedOrder:
			//Slette order
		case network.Cost:
			//Skrive noe
			//Sende til kostnadskanal?
		}		
}
