import "./elevio"

package elevio

import "sync"
import "net"
import "fmt"

type STATES struct {
	IDLE bool
	MOVING bool
	DOOR_OPEN bool
	MOTOR_STOP bool
	INIT bool
}


	door chan bool
	current_floor chan int
	new_order chan int
	elev_dir chan bool
	button_lights chan ButtonEvent
	finished_order chan bool

func initElev(cha)

func fsm_button_pressed()





for{
	select{
	case newOrder := new_order:
		//moving --> update queue
		//idle --> choose direction, setmotordir
	}
}