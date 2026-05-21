package main

import (
	"net/http"
	"os"
	"os/exec"
	"os/user"
)

func shutdown(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("shutdown", "/s", "/t", "0")
	cmd.Run()
}

func info(w http.ResponseWriter, r *http.Request) {
	currentUser, err := user.Current()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	name, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Hostname: " + name + "\nUsername: " + currentUser.Username + "\nUID: " + currentUser.Uid))
}

func main() {
	// conn, err := net.Dial("udp", "8.8.8.8:80")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer conn.Close()

	// localAddr := conn.LocalAddr().(*net.UDPAddr)

	// fmt.Println("Local IP:", localAddr.IP)
	http.HandleFunc("/info", info)
	http.HandleFunc("/shutdown", shutdown)
	http.ListenAndServe(":12500", nil)
}
