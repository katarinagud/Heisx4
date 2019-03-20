import "./elevio"
import "sync"
import "net"
import "fmt"
import "time"





func fsm_run_elev(newOrder <-chan Button, floorReached chan int, orderDone chan<- Button) {

	doorTime := time.NewTimer(3*time.Second)
	var e ElevState

	var dir MotorDirection := ChooseDirection(e)

	var state STATES 
	var order[4][3]int

	for{
		select{
		case newOrder := <- newOrder:

			switch state {
			case IDLE:
				if dir == 0 {
					state = DOOR_OPEN
					elevio.SetDoorOpenLamp(1)
					time.Sleep(3*time.Second)
					elevio.SetDoorOpenLamp(0)
					state = IDLE

				}
				else{
					SetMotorDirection(dir)
					state = MOVING
				}
				
				
			case MOVING:
			
			case DOOR_OPEN:
				if e.Floor == newOrder.Floor {
					elevio.SetDoorOpenLamp(1)
					time.Sleep(3*time.Second)
					elevio.SetDoorOpenLamp(0)
					state = IDLE

				}
			case MOTOR_STOP:
				
			}
		
		case floorArrival := <- floorReached:

			switch state {
			case IDLE:
			
			case MOVING:
				if ShouldStop(e)
					e = ClearAtCurrentFloor(e, func(btn int){ orderDone <- Button{e.Floor, btn} })
					state = DOOR_OPEN
					SetMotorDirection(0)
					doorTime.Reset(3*time.Second)
					elevio.SetDoorOpenLamp(1)
					
					
					// some door shit goes here
					// maybe a timer idk

				
			
			case INIT:
				SetMotorDirection(0)
			}
		case doorTimeOut := <- doorTime.C:
			
			switch state {
			case DOOR_OPEN:
				elevio.SetDoorOpenLamp(0)
				orderDone <- true
			}
		
		}
	}

}





