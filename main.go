package main

import elevio
import queue
import "fmt"
import "./network/bcast"
import "./network/localip"
import "./network/peers"
import "flag"
import "os"

type Button struct {
	Floor int
	Type int
}0


func main(){

    numFloors := 4

	var id string
	flag.StringVar(&id, "id", "", "id of this peer")
	var driver_port string
	flag.StringVar(&driver_port, "driver_port", "", "port to connecto to the elevator")
	flag.Parse()

	if id == "" {
		localIP, err := localip.LocalIP()
		if err != nil {
			fmt.Println(err)
			localIP = "DISCONNECTED"
		}
		id = fmt.Sprintf("peer-%s-%d", localIP, os.Getpid())
	}
	if driver_port == "" {
		driver_port = "15657"
	}

	var all_states map[string]ElevStates

	// Channels

	peerUpdateCh := make(chan peers.PeerUpdate)
	peerTxEnable := make(chan bool)

	buttonTx := make(chan Button)
	buttonRx := make(chan Button)

	drv_buttons := make(chan elevio.ButtonEvent)
	drv_floors  := make(chan int)
	drv_motor_dir := make(chan elevio.MotorDirection)

	local_stateCh := make(chan fsm.ElevStates)
	all_statesCh := make(chan fsm.ElevStates)

	costTx := make(chan queue.CostValue)
	costRx := make(chan queue.CostValue)

	newOrder := make(chan Button)
	orderDone := make(chan bool)

	portCh := make(chan string)

	fsm_move := make(chan bool)



	//Goroutines

	go peers.Transmitter(15647, id, peerTxEnable)
	go peers.Receiver(15647, peerUpdateCh)

	go func(peerUpdateCh <-chan peers.Peerupdate, peerList chan<- []string){
		var p peers.PeerUpdate
		for {
			select {
			case p = <-peerUpdateCh:
			case peerList <- p.Peers:
		}
	}()

	go bcast.Transmitter(16569, buttonTx)
	go bcast.Receiver(16569, 	buttonRx)
	
	go bcast.Transmitter(15000, costTx)
	go bcast.Receiver(15000, costRx)

	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors)

	go elevstates.ElevStates(id, local_state, all_states)

	go queue.UpdateQueue(drv_buttons, all_states, peerUpdateCh, updateOrder)
			
	elevio.Init("localhost:"+driver_port, numFloors)


	go func(){
		var order Order
		for {
			order <- assignedOrder_netRx
			if order.assignedTo == LocalID {
				newOrder <- order
			}
		}
	}()
    
    
    for {
        select {
        case a := <- drv_buttons:
			fmt.Printf("drv_buttons: %#v\n", a)
			buttonTx <- Button{Floor: a.Floor, Type: int(a.Button)}
					
		            
		case a := <- drv_floors:			
			fmt.Printf("drv_floors:  %#v\n", a)
		
			
            
		case p := <-peerUpdateCh:
			fmt.Printf("Peer update:\n")
			fmt.Printf("  Peers:    %q\n", p.Peers)
			fmt.Printf("  New:      %q\n", p.New)
			fmt.Printf("  Lost:     %q\n", p.Lost)

		case a := <-buttonRx:
			fmt.Printf("buttonRx:    %#v\n", a)
		}
    }    
}
