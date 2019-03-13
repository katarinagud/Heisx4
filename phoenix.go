package main

import (
	"encoding/binary"
	. "fmt"
	"log"
	"net"
	"os/exec"
	t "time"
)

var counter uint64
var port = 9999
var buf = make([]byte, 16)

func spawnBackup() {
	(exec.Command("gnome-terminal", "-x", "sh", "-c", "go run phoenix.go")).Run()
	Println("New backup is online!")
}

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	isPrimary := false
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Println("Error, listeningUDP failed!")
	}

	log.Println("I'm the backup!")


	for !(isPrimary) {
		conn.SetReadDeadline(t.Now().Add(2 * t.Second))
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			isPrimary = true
		} else {
			counter = binary.BigEndian.Uint64(buf[:n])
		}
	}
	conn.Close()

	Println("addr: ", addr)
	spawnBackup()
	Println("I'm now the primary!")
	bcastConn, _ := net.DialUDP("udp", nil, addr)


	for {
		if counter%10 == 0 {
			Println("\t*---------------*")
			Println("\t| Number: ", counter, "\t|")
		} else {
			Println("\t| Number: ", counter, "\t|")
		}
		counter++
		binary.BigEndian.PutUint64(buf, counter)
		_, _ = bcastConn.Write(buf)
		t.Sleep(50 * t.Millisecond)
	}
}
