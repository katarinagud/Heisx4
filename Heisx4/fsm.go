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
	Floor := elevio.getFloor()
	Direction int
	State STATES
	Orders [4][3] int
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


func fsm_run_elev(updateOrder <-chan Button, floorReached <-chan int ) {

	for{
		select{
		case newOrder := <- updateOrder:

			switch ElevStates.State {
			case IDLE:
				var dir MotorDirection := ChooseDir(ElevStates.Floor)
				if dir == 0 {
					ElevStates.State = DOOR_OPEN
					elevio.SetDoorOpenLamp(1)
					time.Sleep(3*time.Second)
					elevio.SetDoorOpenLamp(0)
					ElevStates.State = IDLE

				}
				else{
					SetMotorDirection(dir)
					ElevStates.State = MOVING
				}
				
				
			case MOVING:
			
			case DOOR_OPEN:
				if ElevStates.Floor == newOrder.Floor {
					elevio.SetDoorOpenLamp(1)
					time.Sleep(3*time.Second)
					elevio.SetDoorOpenLamp(0)
					ElevStates.State = IDLE

				}
			case MOTOR_STOP:
				
			}
		
		case floorArrival := <- floorReached:

			switch ElevStates.State {
			case IDLE:
			
			case MOVING:
				if ElevStates.Orders[F][] == 
			}
			
		}
	}

}





