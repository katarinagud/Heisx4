
func ChooseDir(floor ElevState) {

}

func ShouldStop(floor_reached int) {
	if ((isOrder(floor_reached) == 1 ) && 
}

func isOrder(floor_reached int) {
	if (ElevState.Orders[floor_reached][0] || ElevState.Orders[floor_reached][1] || ElevState.Orders[floor_reached][2]){
		return true
	}
	else 
		return false
}

