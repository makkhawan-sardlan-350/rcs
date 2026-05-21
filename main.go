package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
)

func shutdown(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("shutdown", "/s", "/t", "0")
	cmd.Run()
}

func main() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	fmt.Println("Local IP:", localAddr.IP)

	http.HandleFunc("/shutdown", shutdown)
	http.ListenAndServe(":12500", nil)
}
