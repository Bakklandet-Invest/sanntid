package network

import (
	"fmt"
	"time"
	"encoding/json"
)

var RecieveChan = make(chan messageUDP)
var sendChan = make(chan messageUDP)

const NotifyInterval = 20*time.Second

func Init(){
	const localListenPort = 13369
	const broadcastListenPort = 13370
	const messageSize = 1024

	err := InitUDP(localListenPort, broadcastListenPort, messageSize, sendChan, RecieveChan)
	if err != nil {
		fmt.Print("InitUDP error: %s \n", err)
	}

	go aliveNotifier()
	go checkOutgoingMessages()
}

func aliveNotifier() {
	alive := Message{Content: Alive, Floor: -1, Button: -1, Cost: -1}
	for {
		MessageCh <- alive
		time.Sleep(NotifyInterval)
	}
}

func checkOutgoingMessages() {
	for {
		fmt.Println("checkOutgoingMessages waiting on MessageCh")
		msg := <-MessageCh

		PrintMessage(msg)

		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			fmt.Print("json.Marshal error: %s \n", err)
		}		
		
		sendChan <- messageUDP{recieveAddr: "broadcast", data: jsonMsg, length: len(jsonMsg)}

		time.Sleep(time.Millisecond)
	}
}

func PrintMessage(msg Message) {
	fmt.Printf("\n-----Message start-----\n")
	switch msg.Content {
	case Alive:
		fmt.Println("I'm alive")
	case NewOrder:
		fmt.Println("New order")
	case CompletedOrder:
		fmt.Println("Completed order")
	case Cost:
		fmt.Println("Cost:")
	default:
		fmt.Println("Invalid message type\n")
	}
	fmt.Printf("Floor: %d\n", msg.Floor)
	fmt.Printf("Button: %d\n", msg.Button)
	fmt.Printf("Cost:   %d\n", msg.Cost)
	fmt.Println("-----Message end-------\n")
}

func ParseMessage(msgUDP messageUDP) Message {
	fmt.Printf("Before parse: %s from %s\n", string(msgUDP.data), msgUDP.recieveAddr)

	var msg Message
	if err := json.Unmarshal(msgUDP.data[:msgUDP.length], &msg); err != nil {
		fmt.Printf("json.Unmarshal error: %s\n", err)
	}

	msg.Addr = msgUDP.recieveAddr
	fmt.Printf("After Unmarshal: %s\n", msg.Addr)
	return msg
}
