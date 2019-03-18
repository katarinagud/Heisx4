import "./elevio"

package elevio

import "sync"
import "net"
import "fmt"


type STATES struct{
	IDLE bool
	MOVING bool
	DOOR_OPEN bool
	MOTOR_STOP bool
	INIT bool
}

type ElevStates struct {
	Floor int
	Direction int
	State STATES
	Orders [4][2] int
}

type ElevQueue struct {
	QueueSystem [4][4]int
	CabCall [4]int
	HallCall [4][2]int
	ID string
}0


func Fsm_update_Elevstates(floor <-chan int, motor_dir <-chan MotorDirection, updateOrder <-chan Button) {

	for{
		select{
	case a := <- floor:
		ElevStates.Floor = a
		ElevStates.States = State
	case a := <- motor_dir:
		ElevStates.Direction = a
		ElevStates.States = State
	case a:= <- updateOrder:
		if a.Type != 2{
			ElevStates.Orders[a.Floor][a.Type] = 1
			ElevStates.States = State
		}

			
	}
	

	}
}


func fsm_button_pressed()





for{
	select{
	case newOrder := new_order:
		//moving --> update queue
		//idle --> choose direction, setmotordir
	}
}