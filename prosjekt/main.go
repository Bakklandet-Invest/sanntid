package main

import(
	"network"
	"time"
	"fmt"
	"strconv"
	"net"
	."control"
	."driver"
	"filehandler"
)

var liftsOnline = make(map[string]network.ConnectionUDP)
var LiftsOnlineInfo = make(map[string]network.ElevatorInfo)
var disconnElevChan = make(chan network.ConnectionUDP)
var myID string = findMyID()
var uncompleteOrders Matrix

func selectMaster(lifts map[string]network.ConnectionUDP, checkMasterChan chan string)string {
	master := myID
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

func main(){
	hang := make(chan int)

	updateOutChan := make(chan network.ElevatorInfo)
	updateInChan := make(chan network.ElevatorInfo)
	checkMasterChan := make(chan string)
	newOrderChan := make(chan ButtonSignal,20)
	completeOrderChan := make(chan ButtonSignal)
	extOrderChan := make(chan ButtonSignal)
	fromMasterChan := make(chan ButtonSignal)

	updateLightsSlaveChan := make(chan int)
	completeOrderBroadcastChan := make(chan ButtonSignal)	
	
	// Backup works, but is not sent to InitElevator due to limited time
	backupChan := make(chan Matrix)
	go backupHandler(backupChan)
	
	go networkHandler(updateInChan, checkMasterChan, newOrderChan, completeOrderChan, completeOrderBroadcastChan, updateLightsSlaveChan)
	go InitElevator(updateOutChan,  checkMasterChan, completeOrderChan, extOrderChan, fromMasterChan)
	go Master(updateOutChan, checkMasterChan, newOrderChan, completeOrderChan, extOrderChan, fromMasterChan, completeOrderBroadcastChan, updateLightsSlaveChan)

	<-hang
}

func Slave(updateOutChan chan network.ElevatorInfo, checkMasterChan chan string, newOrderChan chan ButtonSignal, completeOrderChan chan ButtonSignal, extOrderChan chan ButtonSignal, fromMasterChan chan ButtonSignal, completeOrderBroadcastChan chan ButtonSignal, updateLightsSlaveChan chan int){		
	var masterAddr string 
	var master string
	fmt.Println("Now running as slave")
	for{
		select{
			case master = <- checkMasterChan:
				if master == myID{
					go Master(updateOutChan, checkMasterChan, newOrderChan, completeOrderChan, extOrderChan, fromMasterChan, completeOrderBroadcastChan, updateLightsSlaveChan)
					return
				}
				masterAddr = liftsOnline[master].Addr
			case update := <- updateOutChan:
				updateMsg := network.Message{Content: network.Info, Addr: "broadcast", ElevInfo: update}
				network.MessageChan <- updateMsg
			case extOrdButtonSignal := <- extOrderChan:
				extOrdMsg := network.Message{Content: network.NewOrder, Addr: masterAddr, Floor: extOrdButtonSignal.Floor, Button: extOrdButtonSignal.Button}
				network.MessageChan <- extOrdMsg 
			case orderButtonSignal := <- newOrderChan:
				fromMasterChan <- orderButtonSignal
			case completeOrder := <- completeOrderChan:
				complOrdMsg := network.Message{Content: network.CompletedOrder, Addr: masterAddr, Floor: completeOrder.Floor, Button: completeOrder.Button}
				network.MessageChan <- complOrdMsg
			case <-updateLightsSlaveChan:
				for i := 0; i < network.NumFloors; i++{
					for j := 0; j < network.NumButtons-1; j++{
						if LiftsOnlineInfo[master].Matrix[i][j] {
							Elev_set_button_lamp(ButtonSignal{Light: 1})
						} else {
							Elev_set_button_lamp(ButtonSignal{Light: 0})						
						}
					}	
				}
		}	
	}
}

func Master(updateOutChan chan network.ElevatorInfo, checkMasterChan chan string, newOrderChan chan ButtonSignal, completeOrderChan chan ButtonSignal, extOrderChan chan ButtonSignal, fromMasterChan chan ButtonSignal, completeOrderBroadcastChan chan ButtonSignal, updateLightsSlaveChan chan int){
	fmt.Println("Now running as master")	
	for{
		select{
			case master := <- checkMasterChan:
				if master != myID{
					go Slave(updateOutChan, checkMasterChan, newOrderChan, completeOrderChan, extOrderChan, fromMasterChan, completeOrderBroadcastChan, updateLightsSlaveChan)
					return
				}
			case update := <- updateOutChan:
				updateMsg := network.Message{Content: network.Info, Addr: "broadcast", ElevInfo: update}
				network.MessageChan <- updateMsg
			case orderButtonSignal := <- newOrderChan: //Finds best elevator for job
				if uncompleteOrders[orderButtonSignal.Floor][orderButtonSignal.Button] == true{
					continue
				} 
				uncompleteOrders[orderButtonSignal.Floor][orderButtonSignal.Button] = true
				Elev_set_button_lamp(orderButtonSignal)
				
				heisID := myID
				cost := 100
				newCost := 101				
				for key := range LiftsOnlineInfo {
					heisInfo := LiftsOnlineInfo[key]					
					if heisInfo.Dir == 0 {
						newCost = SimpleCost(heisInfo.Floor, orderButtonSignal.Floor)			
					} else if (heisInfo.Dir > 0 && ((heisInfo.Floor < orderButtonSignal.Floor) && orderButtonSignal.Button == BUTTON_CALL_UP)){
						newCost = SimpleCost(heisInfo.Floor, orderButtonSignal.Floor)
					} else if (heisInfo.Dir < 0 && ((heisInfo.Floor > orderButtonSignal.Floor) && orderButtonSignal.Button == BUTTON_CALL_DOWN)){
						newCost = SimpleCost(heisInfo.Floor, orderButtonSignal.Floor)
					} else {
						newCost = ComplexCost(heisInfo.Dir, heisInfo.Floor, heisInfo.Matrix, orderButtonSignal.Floor)
					}
					if newCost < cost {
						cost = newCost
						heisID = key
					}
				}
				if heisID == myID{
					fromMasterChan<-orderButtonSignal
				}else{
					addrOrderReciever := liftsOnline[heisID].Addr
					newOrderMsg := network.Message{Content: network.NewOrder, Addr: addrOrderReciever, Floor: orderButtonSignal.Floor, Button: orderButtonSignal.Button}
					network.MessageChan <- newOrderMsg
				}
				updateMsg := network.Message{Content: network.Light, Addr: "broadcast", ElevInfo: network.ElevatorInfo{Matrix: uncompleteOrders}}
				network.MessageChan <- updateMsg
			case extOrdButtonSignal := <- extOrderChan:
				newOrderChan <- extOrdButtonSignal
			case completeOrderLocal := <- completeOrderChan: 
				uncompleteOrders[completeOrderLocal.Floor][completeOrderLocal.Button] = false
				Elev_set_button_lamp(completeOrderLocal)
			case completeOrderBroadcast := <-completeOrderBroadcastChan:
				uncompleteOrders[completeOrderBroadcast.Floor][completeOrderBroadcast.Button] = false
				Elev_set_button_lamp(completeOrderBroadcast)
			case <-updateLightsSlaveChan:
				continue
		}
	}
}

func findMyID() string{
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                IDstr := ipnet.IP.String()
                IDstr = IDstr[12:]
                return IDstr//ID
            }
        }
	} 
	return "69"
}

func backupHandler(backupChan chan Matrix){
	loadBackupTime := time.Second*10
	saveBackupTime := time.Second
	loadTimer := time.NewTimer(loadBackupTime)
	saveTimer := time.NewTimer(saveBackupTime)
	for{
		select{
		case <-saveTimer.C: 
			filehandler.SaveBackup(LiftsOnlineInfo)
			saveTimer.Reset(saveBackupTime)
		case conn := <- disconnElevChan:
			id := network.FindID(conn.Addr)
			m := matrixCompareOr(LiftsOnlineInfo[id].Matrix, LiftsOnlineInfo[myID].Matrix)
			backupChan<-m 
			delete(LiftsOnlineInfo, id)
		case <-loadTimer.C:
			backupMap := filehandler.LoadBackup()
			var backupMatrix /*(network.ElevatorInfo).*/Matrix = backupMap[myID].Matrix
			if backupMatrix != LiftsOnlineInfo[myID].Matrix{
				backupMatrix = matrixCompareOr(backupMatrix, LiftsOnlineInfo[myID].Matrix)
				backupChan<-backupMatrix
			}
			loadTimer.Reset(loadBackupTime)
		}
	}
}

func matrixCompareOr(m1 Matrix, m2 Matrix) (Matrix){
	var m3 Matrix
	for i := 0; i < network.NumFloors; i++ {
		for j := 0; j < network.NumButtons; j++{
			if m1[i][j] == true || m2[i][j] == true{
				m3[i][j] = true
			}
		}
	}
	return m3
}

func networkHandler(updateInChan chan network.ElevatorInfo, checkMasterChan chan string, newOrderChan chan ButtonSignal, completeOrderChan chan ButtonSignal, completeOrderBroadcastChan chan ButtonSignal, updateLightsSlaveChan chan int){
	network.Init()
	for{
		select{
		case msg := <-network.RecieveChan:
			messageHandler(network.ParseMessage(msg), updateInChan, checkMasterChan, newOrderChan, completeOrderChan, completeOrderBroadcastChan, updateLightsSlaveChan)
		case conn := <- disconnElevChan:
			deleteLift(conn.Addr, checkMasterChan)
		}
			
			
	}
}

func messageHandler(msg network.Message, updateInChan chan network.ElevatorInfo, checkMasterChan chan string, newOrderChan chan ButtonSignal, completeOrderChan chan ButtonSignal, completeOrderBroadcastChan chan ButtonSignal, updateLightsSlaveChan chan int) {
	id := network.FindID(msg.Addr)
	switch msg.Content{
		case network.Alive:
			if conn, alive := liftsOnline[id]; alive{
				conn.Timer.Reset(network.ResetConnTime)
			} else{
				newConn := network.ConnectionUDP{msg.Addr, time.NewTimer(network.ResetConnTime)}
				
				liftsOnline[id] = newConn
				go connTimer(newConn)
				selectMaster(liftsOnline, checkMasterChan) 
			}
		case network.NewOrder:
			newOrderChan<-ButtonSignal{Floor: msg.Floor, Button: msg.Button, Light: 1} 
		case network.CompletedOrder:
			completeOrderBroadcastChan <- ButtonSignal{Floor: msg.Floor, Button: msg.Button}
		case network.Info:
			LiftsOnlineInfo[id] = msg.ElevInfo
			updateLightsSlaveChan<-1
		case network.Light:
			updateLightsSlaveChan<-1
		}		
}

func connTimer(conn network.ConnectionUDP){
	for{
	<-conn.Timer.C
	disconnElevChan <- conn
	}
}

func deleteLift(addr string, checkMasterChan chan string){
	id := network.FindID(addr)
	delete(liftsOnline, id)
	selectMaster(liftsOnline, checkMasterChan) 
}
