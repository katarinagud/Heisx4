

import "./elevio"
import "fmt"
import "math/rand"
import

type ElevQueue struct {
	QueueSystem [4][4]int
	CabCall [4]int
	HallCall [4][2]int
	ID string
}


func UpdateQueue(buttonPressed <-chan Button, all_states <-chan map[string]fsm.ElevStates, peers <-chan peers.PeerUpdate, updateOrder chan<- Button){
	for{
		select{
		case a := <- buttonPressed:
			if a.Type != 2 {
				ElevQueue.HallCall[a.Floor][a.Type] = 1

				var costs map[string][4] int
				var cost int
				for id, elev_state := range all_states{
					cost = calculateCost(elev_state)
					costs[id] = cost
					}
				//CALCULATE COST and compare

				if lowest cost
				//update fsm.ElevState
				//updateOrder
			}
			else{
				ElevQueue.CabCall[a.Floor] = 1
			}
		case a := <- all_states:
			
		
		}
	}
}

func calculateCost(elev_state fsm.ElevState){

}







