package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// func writeToScocket(conn *websocket.Conn, msg []byte) {
// 	if err = conn.WriteMessage(59875, msg);
// 	err != nil {
// 		return
// 	}
// }
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/index.html")
	})

	// http.HandleFunc("/agent", func(w http.ResponseWriter, r *http.Request){
	// 	http.re
	// })

	http.HandleFunc("/stdout", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		msgType, msg, _ := conn.ReadMessage()

		_ = conn.WriteMessage(msgType, msg)
		time.Sleep(3 * time.Second)

		env := "/Users/isaraovin/.local/share/virtualenvs/Scraping-R21sHmfA/bin/python"
		script := "/Users/isaraovin/Documents/Projects/LogMonitor/dummy.py"

		cmd := exec.Command(env, script)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Start()

		buf := bufio.NewReader(stdout) // Notice that this is not in a loop
		num := 1
		for {
			line, _, _ := buf.ReadLine()
			is_logging := len(line)
			if is_logging == 0 {
				break
			}
			num += 1
			fmt.Println(string(line))
			if err = conn.WriteMessage(msgType, line); err != nil {
				return
			}
		}

	})

	http.ListenAndServe(":8080", nil)
}
