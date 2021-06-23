package main

import (
  "log"
  "fmt"
  "time"

  "golang.org/x/net/websocket"
)

var (
  locate = "localhost:8080/test"
  origin = "http://" + locate
  ws_url = "ws://" + locate + "/ws"
)

// struct for message
type EchoMsg struct {
  Msg string `json: msg`
}

func main() {
  ws, err := websocket.Dial(ws_url, "", origin)
  if err != nil {
    log.Fatal(err)
  }

  go receiveMsg(ws)

  sendMsg(ws, "Hi.")
  sendMsg(ws, "I'm client.")

  time.Sleep(1*time.Second)
  _ = ws.Close()

  defer log.Printf("End Webscoket.")
}


func sendMsg(ws *websocket.Conn, msg string) {
    // var sndMsg = EchoMsg{msg}

    websocket.Message.Send(ws, msg)
    fmt.Printf("Send data=%#v\n", msg)
}

func receiveMsg(ws *websocket.Conn) {
    // var rcvMsg EchoMsg
    var msg string

    for {
        websocket.Message.Receive(ws, &msg)
        fmt.Printf("Receive data=%#v\n", msg)
    }
}

