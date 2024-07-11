// package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}

// env := "/Users/isaraovin/.local/share/virtualenvs/Scraping-R21sHmfA/bin/python"
// script := "/Users/isaraovin/Documents/Projects/LogMonitor/dummy.py"

// cmd := exec.Command(env, script)
// stdout, err := cmd.StdoutPipe()
// if err != nil {
// 	log.Fatal(err)
// }
// cmd.Start()

// buf := bufio.NewReader(stdout) // Notice that this is not in a loop
// num := 1
// for {
// 	line, _, _ := buf.ReadLine()
// 	is_logging := len(line)
// 	if is_logging == 0 {
// 		os.Exit(0)
// 	}
// 	num += 1
// 	fmt.Println(string(line))
// }
