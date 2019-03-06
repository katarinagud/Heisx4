

import "./elevio"
import "fmt"
import "math/rand"

type ElevQueue struct {
	QueueSystem [4][4]int
	CabCall [4]int
	HallCall [4][2]int
	ID string
}

func UpdateQueueFromIO(button ButtonEvent) {
	if button.Button == 2 {
		ElevQueue.CabCall[button.Floor] == 1
	}

	if button.Button != 2 {
		ElevQueue.HallCall[button.Floor][button.Button] == 1
	}

	return ElevQueue
}

func UpdateQueueFromNetwork(button <-chan Button) {
	if button.Type == 2 {
		ElevQueue.CabCall[button.Floor] == 1
	}

	if button.Type != 2 {
		ElevQueue.HallCall[button.Floor][button.Type] == 1
	}

	cost int = rand.Intn(100)

	return cost
}





