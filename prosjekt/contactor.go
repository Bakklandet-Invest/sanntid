package main

import(
	"network"
	"time"
	"fmt"
	"strconv"
	"net"
	."control"
)

var liftsOnline = make(map[string]network.ConnectionUDP)
var disconnElevChan = make(chan network.ConnectionUDP)
var myID string = FindElevID()

func writemap() {
	for {
		fmt.Println("Size av liftsOnline:",len(liftsOnline))
		fmt.Println("liftsOnline:",liftsOnline)
		fmt.Println("Master har ID:",selectMaster(liftsOnline, checkMasterChan))
		time.Sleep(time.Second*10)
	}
}

func selectMaster(lifts map[string]network.ConnectionUDP, checkMasterChan chan string)string {
	master := "-1"
	for key := range lifts{
		m, _ := strconv.Atoi(master)
		k, _ := strconv.Atoi(key)
		if k > m{
			master = key
		}
	}
	checkMasterChan<-master
	return master
}

//func masterAlive()bool{
//}
/* ---HUSKELISTE
	ElevInfo må pakkes i ElevatorRun
	* ConnTimer tar kun inn network.ConnectionUDP, og vet ikke ID'en
	* -> fikset at den sletter nå
*/ /* */
func main(){
	updateOutChan := make(chan network.Message)
	updateInChan := make(chan network.Message)//Trenger ikke egen case, messageHandler tar ansvar
	checkMasterChan := make(chan string)	
	newOrderChan := make(chan ButtonSignal)
	completeOrderChan := make(chan ButtonSignal)
	extOrderChan := make(chan ButtonSignal)
	fromMasterChan := make(chan ButtonSignal)//Kanalen som går inn til elev og gir oppdrag
	
	// Holder orden på master/slave-rollen	
	//terminateChan := make(chan bool)
	//terminatedChan := make(chan bool)

	
	
	go networkHandler(updateInChan, checkMasterChan, newOrderChan, completeOrderChan)
	go InitElevator(updateOutChan, updateInChan, checkMasterChan, newOrderChan, completeOrderChan, extOrderChan, fromMasterChan)
	go Master(updateOutChan, checkMasterChan, newOrderChan, completeOrderChan, extOrderChan, fromMasterChan)
	
}

//func checkMaster(checkMasterChan chan string){


func Slave(updateOutChan chan network.Message, checkMasterChan chan string, newOrderChan chan ButtonSignal, completeOrderChan chan ButtonSignal, extOrderChan chan ButtonSignal, fromMasterChan chan ButtonSignal){		
	//go networkHandler()
	var masterAddr string //evt lagre sin egen slik at meldingene sendes til seg selv
	// og ikke forsvinner i vente på en master
	for{
		select{
			case master := <- checkMasterChan:
				if master == myID{
					terminateChan<-true
					<-terminatedChan
					go Master(updateOutChan, checkMasterChan, newOrderChan, completeOrderChan, extOrderChan, fromMasterChan)
					return
				} //PROBLEM UNDER- i init fasen, vil ikke masterADdr eksistere
				masterAddr = liftsOnline(master)
			case update := <- updateOutChan:
				updateMsg := network.Message{Content: network.Info, Addr: "broadcast", ElevInfo: update}
				network.MessageChan <- updateMsg
			case extOrdButtonSignal := <- extOrderChan:
				extOrdMsg := network.Message{Content: network.NewOrder, Addr: <addresse til master>, Floor: extOrdButtonSignal.Floor, Button: extOrdButtonSignal.Button}
				network.MessageChan <- extOrdMsg //Få pakket forunftig beskjed her først da, med recieverAddr til MASTER
			// PROBLEM - lagre extOrd lokalt i tilfelle intet svar, kjøre da for å være sikker??
	
			case orderButtonSignal := <- newOrderChan:
				fromMasterChan <- orderButtonSignal
			case completeOrder := <- completeOrderChan:
				complOrdMsg := network.Message{Content: network.CompletedOrder, Addr: <addresse til master>, Floor: extOrdButtonSignal.Floor, Button: extOrdButtonSignal.Button}
				network.MessageChan <- complOrdMsg
		}	
	}
}

func Master(updateOutChan chan network.Message, checkMasterChan chan string, newOrderChan chan ButtonSignal, completeOrderChan chan ButtonSignal, extOrderChan chan ButtonSignal, fromMasterChan chan ButtonSignal){

	for{
		select{
			case master := <- checkMasterChan:
				if master != myID{
					terminateChan<-true
					<-terminatedChan
					go Slave(updateOutChan, checkMasterChan, newOrderChan, completeOrderChan, extOrderChan, fromMasterChan)
					return
				}
			case update := <- updateOutChan:
				updateMsg := network.Message{Content: network.Info, Addr: "broadcast", ElevInfo: update}
				network.MessageChan <- updateMsg
			case extOrdButtonSignal := <- extOrderChan:
				extOrdMsg := network.Message{Content: network.NewOrder, Addr: <addresse til master>, Floor: extOrdButtonSignal.Floor, Button: extOrdButtonSignal.Button}
				network.MessageChan <- extOrdMsg //Få pakket forunftig beskjed her først da, med recieverAddr til MASTER
			case orderButtonSignal := <- newOrderChan: //Finner beste heis til å ta jobben og sender til den
				//returnerer heis best egnet for jobbet
				heisID = costfunction(orderButtonSignal) //order inneholder opp/ned+etasje
		// --- LAGRE ordren i uncompleteOrders og slette igjen når completed er mottatt??
				if heisID == myID{
					fromMasterChan<-orderButtonSignal
				}else{
					addrOrderReciever := liftsOnline(heisID).Addr
					newOrderMsg := network.Message{Content: network.NewOrder, Addr: addrOrderReciever, Floor: orderButtonSignal.Floor, Button: orderButtonSignal.Button}
					network.MessageChan <- newOrderMsg
				}
			case completeOrder := <- completeOrderChan:
				// slette fra uncompleteOrders (PS; SYNKE uncompleteOrders til alle for backup?) (no orders lost)
				
		}
	}
}

func FindElevID() string/*int*/ {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                IDstr := ipnet.IP.String()
                IDstr = IDstr[12:]
                //ID, _ := strconv.Atoi(IDstr)
                return IDstr//ID
            }
        }
    }else{
		fmt.Println(err)
		return "69"
	} 
}

/*
func main(){
	var kanal = make(chan int)
	
	go networkHandler()
	
	<-kanal
}*/

func networkHandler(updateInChan chan network.Message, checkMasterChan chan string, newOrderChan chan ButtonSignal, completeOrderChan chan ButtonSignal){
	network.Init()
	go writemap()
	for{
		select{
		case msg := <-network.RecieveChan:
			messageHandler(network.ParseMessage(msg), updateInChan, checkMasterChan, newOrder, completeOrderChan)
		case conn := <- disconnElevChan:
			deleteLift(conn.Addr)
		}
			
			
	}
}

//  msg.Addr byttet med id i liftsonline(key)
func messageHandler(msg network.Message, updateInChan chan network.Message, checkMasterChan chan string, newOrderChan chan ButtonSignal, completeOrderChan chan ButtonSignal,) {
	id := network.FindID(msg.Addr)
	switch msg.Content{
		case network.Alive:
			if conn, alive := liftsOnline[id]; alive{
				conn.Timer.Reset(network.ResetConnTime)
			} else{
				newConn := network.ConnectionUDP{msg.Addr, time.NewTimer(network.ResetConnTime)}
				
				liftsOnline[id] = newConn
				go connTimer(newConn)
				selectMaster(liftsOnline, checkMasterChan) //CHECKMASTER
			}
		case network.NewOrder:
			//kanskje skrive ut noe?
			// KANAL skrive orderen til canalen
			/*tempButtonSignal := ButtonSignal{Floor: msg.Floor, Button: msg.Button} 
			newOrderChan<-tempButtonSignal*/
			newOrderChan<-ButtonSignal{Floor: msg.Floor, Button: msg.Button} 
			//cost :=  69 //costfunksjonen inn her
			
			//costMsg := network.Message{Content: network.Cost, Floor: msg.Floor, Button: msg.Button, Cost: cost}
			//network.MessageChan <- costMsg
		case network.CompletedOrder:
			//Slette order - denne casen er kun for master??
			completeOrderChan <- ButtonSignal{Floor: msg.Floor, Button: msg.Button}
		case network.Info:
			// LAGRER INFO OM HEISEN M/ TILHØRENDE ID, hvor?
// bruke updateInChan			
		}		
}

func connTimer(conn network.ConnectionUDP){
	for{
	<-conn.Timer.C
	disconnElevChan <- conn
	}
}

//endret slik at den sletter vha idkey, ikke addr key
func deleteLift(addr string){
	id := network.FindID(addr)
	delete(liftsOnline, id)
	selectMaster(liftsOnline, checkMasterChan) //CHECKMASTER
}
