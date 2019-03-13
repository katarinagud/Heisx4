

import "./elevio"
import "fmt"
import "math/rand"
import

type ElevQueue struct {
	QueueSystem [4][4]int
	CabCall [4]int
	HallCall [4][2]int
	ID string
}0

type CostValue struct {
	cost int
	port string
}

func UpdateQueue(buttonPressed chan Button, all_states chan map[string]fsm.ElevState, order_accepted chan Button){

}

func GetPortNumb(port <-chan portCh) {
	CostValue.port = <- port
}

func SendCost(CostCh chan<- CostValue) {
	CostCh <- CostValue
}







