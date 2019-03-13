package main

import "./elevio"
import "./queue"
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

	// Channels

	peerUpdateCh := make(chan peers.PeerUpdate)
	peerTxEnable := make(chan bool)

	buttonTx := make(chan Button)
	buttonRx := make(chan Button)

	drv_buttons := make(chan elevio.ButtonEvent)
	drv_floors  := make(chan int)

	local_state := make(chan fsm.ElevState)
	all_states := make(chan fsm.ElevState)

	costTx := make(chan queue.CostValue)
	costRx := make(chan queue.CostValue)

	order_accepted := make(chan queue.ElevQueue)

	portCh := make(chan string)

	fsm_move := make(chan bool)



	//Goroutines

	go peers.Transmitter(15647, id, peerTxEnable)
	go peers.Receiver(15647, peerUpdateCh)

	go bcast.Transmitter(16569, buttonTx)
	go bcast.Receiver(16569, 	buttonRx)
	
	go bcast.Transmitter(15000, costTx)
	go bcast.Receiver(15000, costRx)

	go elevio.PollButtons(drv_buttons)
	go elevio.PollFloorSensor(drv_floors)

	go elevstates.ElevStates(id, local_state, all_states)

	go queue.UpdateQueue(buttonRx, all_states, order_accepted)	

	go fsm.SetDir(fsm_move)

	go timer.""()

	elevio.Init("localhost:"+driver_port, numFloors)


	
    
    
    for {
        select {
        case a := <- drv_buttons:
			fmt.Printf("drv_buttons: %#v\n", a)
			buttonTx <- Button{Floor: a.Floor, Type: int(a.Button)}
					
		            
		case a := <- drv_floors:			
			fmt.Printf("drv_floors:  %#v\n", a)
		
			
		//case b := <-orderAccepted:
			// Bestemme hvilken av heisene som skal ta orderen
            
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

//func updateQueue(oldQueue){
	//legge ordre i egen matrise
	//sende ut ny ordrematrise
	//vente på godkjenning fra alle heiser}