//Program for å lytte

package main
import "net"
import "fmt"
import "log"

func listenAddr(udpListenChan chan<-int, finished chan bool){
    udpAddr, _ := net.ResolveUDPAddr("udp", ":30000")

    udpConn, _ := net.ListenUDP("udp", udpAddr)

    buffer := make([]byte, 1024)

    n, _, _ := udpConn.ReadFromUDP(buffer)

    fmt.Println(string(buffer[0:n]))

    finished <- true
        return
}

func writeAddr(finished chan bool){

    udpAddr, _ := net.ResolveUDPAddr("udp", "10.100.23.242:30000")

    conn, _ := net.DialUDP("udp", nil, udpAddr)

    message := []byte("Hello from gr 20!!")

    _, err := conn.Write(message)

    if err != nil {
        log.Println(err)
    }

    finished <- true
        return
    
    defer conn.Close()


}


func main(){

    // udpWriteChan := make(chan int)
    // udpListenChan := make(chan int)
    // quit := make(chan bool)
    finished := make(chan bool)

    //go listenAddr(udpListenChan, finished)
    go writeAddr(finished)

    <-finished 
    <-finished


    // ln, err := net.Listen("udp", ":30000")
    // conn, _ := net.DialUDP("udp", nil, &net.UDPAddr{IP:[]byte{127,0,0,1}, Port:20000, Zone:"")
    // defer conn.Close()
    // conn.Write([]byte("hello gruppe20"))
}