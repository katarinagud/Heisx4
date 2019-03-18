package main

type T struct {
	State ElevStates
	ID string
}


func ElevStates(local_id string, local_state ElevStates, all_statesCh chan<- map[string]ElevStates ){

    var states map[string]ElevStates

	netSend := make(chan T)
	netRecv := make(chan T)

    go bcast.Transmitter(15001, netSend)
	go bcast.Receiver(15001, netRecv)

    for{
        select{
		case a := <- local_state:
				netSend <- T{a, local_id}
				states[local_id] = a
		case a := <- netRecv:
				if a.ID != local_id {
					states[a.ID] = a.State
				}
				
			
		case all_statesCh <- states:

	

        }
    }
}