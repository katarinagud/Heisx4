package main

import (
	"../conn"
	"fmt"
	"net"
	"sort"
	"time"
) 

type T struct {
	State ElevStates
	ID string
}


func ElevStates(local_id string, local_state ElevStates, all_states map[string]ElevStates ){

	var states map[string]ElevStates	

	netSend := make(chan T)
	netRecv := make(chan T)

    go bcast.Transmitter(15001, netSend)
	go bcast.Receiver(15001, netRecv)

    for{
        select{
		case a := <- local_state:
				states[local_id] = a
		case a := <- netRecv:
				if a.ID != local_id {
					states[a.ID] = a.State
				}
			case a := <-ticker
				//send state periodically on net
				netSend <- T{a, local_id}
				// do lights here: hall for all elevs, cab for us 
				
			
		case all_states <- states:

	

		}
		
    }
}