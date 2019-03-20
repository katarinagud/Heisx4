
import ""./network"
import "./elevio"
import "types"


func ChooseDirection(e ElevState) {
    switch(e.Direction){
	case MD_Up:
		if requests_above(e)
			return MD_Up
		else if requests_below(e)
			return MD_Down
		else 
		return MD_Stop

	case MD_Down:
    case MD_Stop: // there should only be one request in this case. Checking up or down first is arbitrary.
		if requests_below(e)
			return MD_Down
		else if requests_above(e)
			return MD_Up
		else 
			return MD_Stop
    default:
        return MD_Stop
}
}

func requests_above(Elevator e)bool{
    for (f int := e.Floor+1; f < N_FLOORS; f++){
        for (btn int:= 0; btn < N_BUTTONS; btn++){
            if (e.Orders[f][btn]){
                return true;
            }
        }
    }
    return 0;
}


func requests_below(Elevator e) bool{
    for (f int := 0; f < e.Floor; f++){
        for (btn int := 0; btn < N_BUTTONS; btn++){
            if (e.Orders[f][btn]){
                return true;
            }
        }
    }
    return 0;
}


func ShouldStop(e ElevState) bool {
    switch (e.Direction){
    case MD_Down:
        return
            (e.Orders[e.Floor][BT_HallDown] ||
            e.Orders[e.Floor][BT_Cab]      ||
            !requests_below(e))
    case MD_Up:
        return
            (e.Orders[e.Floor][BT_HallUp]   ||
            e.Orders[e.Floor][BT_Cab]      ||
            !requests_above(e))
    case MD_Stop:
    default:
        return 1
	}
}

func ClearAtCurrentFloor(e ElevState, onClearedOrder func(btnType int)) ElevState {
	for btn := 0; btn <= 2; btn++ {
		if e.Orders[e.Floor][btn] == 2 {
			e.Orders[e.Floor][btn] = 0
			onClearedOrder(btn)
		}
	}
}

