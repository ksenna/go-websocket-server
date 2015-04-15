package main

import (
  "code.google.com/p/go.net/websocket"
  "fmt"
  "log"
  "net/http"
)

func wsRepeat(ws *websocket.Conn) {
  var err error

  for {
    var reply string

    if err = websocket.Message.Receive(ws, &reply); err != nil {
      fmt.Println("Cannot receive ws message")
      break
    }

    fmt.Println("Received message back from ws client: " + reply)

    msg := "Received: " + reply
    fmt.Println("Sending message to ws client: " + msg)

    if err = websocket.Message.Send(ws, msg); err != nil {
      fmt.Println("Cannot send ws message")
      break
    }
  }
}

func main() {
  http.Handle("/", websocket.Handler(wsRepeat))

  if err := http.ListenAndServe(":1234", nil); err != nil {
    log.Fatal("ListenAndServe", err)
  }
}
