package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		fmt.Println("on connection:", so.ID())
		return nil
	})

	server.OnEvent("/chat","chat",func(so socketio.Conn, msg string) string {
		so.SetContext(msg)
		fmt.Println("chat:", msg)
		return "recv" + msg
	})


	server.OnError("/", func(so socketio.Conn, err error) {
		fmt.Println("error:", err)
	})

	server.OnDisconnect("/", func(so socketio.Conn, reason string) {
		fmt.Println("closed:", reason)
	})

	go server.Serve()
	defer server.Close()
	
	http.Handle("/socket.io/", server)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Serving as localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}