package main

import "./elevio"
import "fmt"
import "math/rand"
import "./types"

type ElevQueue struct {
	QueueSystem [4][4]int
	CabCall [4]int
	HallCall [4][2]int
	ID string
}




func Assigner(buttonPressed <-chan elevio.ButtonEvent, all_states <-chan map[string]fsm.ElevState, peerList <-chan []string, assignedOrder chan<- Order){
	for{
		select{
		case a := <- buttonPressed:
			states := <-all_states

			// TODO: filter out dead peers via peerList
			bestID := findBest(a, states)
			assignedOrder <- AssignedOrder{a.Floor, a.ButtonType, bestID}		
		}
	}
}

func findBest(btn elevio.ButtonEvent, states map[string]fsm.ElevState) string {
	bestCost := int.max
	bestID := LocalID
	for id, state := range(states) {
		state_cpy := state	// copy necessary??
		state_cpy.Orders[btn.Floor][btn.ButtonType] = 2
		c := timeToIdle(state_cpy)
		if c < bestCost {
			bestCost = c
			bestID = id
		}
	}
	return bestID
}


func timeToIdle(state fsm.ElevState) int {
	const travelTime = 2500
	const doorOpenTime = 3000
    duration := 0
    
    switch state.STATE {
    case IDLE:
        state.Direction = requests_chooseDirection(e)
        if(state.Direction == MD_Stop){
            return duration;
        }
        break
    case MOVING:
        duration += travelTime/2
        state.Floor += state.Direction
        break
    case DOOR_OPEN:
        duration -= doorOpenTime/2
    }


    for {
        if(ShouldStop(state)){
            state = ClearAtCurrentFloor(state, NULL)
            duration += doorOpenTime
            state.Direction = ChooseDirection(state)
            if(state.Direction == MD_Stop){
                return duration
            }
        }
        state.Floor += state.direction
        duration += travelTime
    }
}






