package main
import "net"
import "log"



func main() {

	tcpAddr, _ := net.ResolveTCPAddr("tcp", "10.100.23.242:33546")



	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	defer conn.Close()

	message := []byte("Hello from gr 20!!\000")

    _, err := conn.Write(message)

    if err != nil {
        log.Println(err)
    }


}