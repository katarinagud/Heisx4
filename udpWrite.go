package main
import "net"
import "log"





// func writeAddr(){

//     udpAddr, _ := net.ResolveUDPAddr("udp", "10.100.23.242:30000")

// 	conn, _ := net.DialUDP("udp", nil, udpAddr)
	
// 	defer conn.Close()

//     message := []byte("Hello from gr 20!!")

//     _, err := conn.Write(message)

//     if err != nil {
//         log.Println(err)
//     }

    
// }

func main() {

	udpAddr, _ := net.ResolveUDPAddr("udp", "10.100.23.242:20001")

	conn, _ := net.DialUDP("udp", nil, udpAddr)
	
	defer conn.Close()

    message := []byte("Hello from gr 20!!")

    _, err := conn.Write(message)

    if err != nil {
        log.Println(err)
    }


}