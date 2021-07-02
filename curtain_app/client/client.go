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
  sendMsg(ws, "Bye.")

  time.Sleep(10*time.Second)
  _ = ws.Close()

  defer log.Printf("End Webscoket.\n")
}


func sendMsg(ws *websocket.Conn, msg string) {
  // var sndMsg = EchoMsg{msg}

  err := websocket.Message.Send(ws, msg)
  if err != nil {
    log.Print(err)
  } else {
    fmt.Printf("Send data=%#v\n", msg)
  }
}

func receiveMsg(ws *websocket.Conn) {
  // var rcvMsg EchoMsg
  var msg string

  for err := websocket.Message.Receive(ws, &msg); err == nil; err = websocket.Message.Receive(ws, &msg) {
    fmt.Printf("Receive data=%#v\n", msg)
  }
  defer log.Printf("End Receiving.\n")

  /*
  for {
    err := websocket.Message.Receive(ws, &msg)
    if err != nil {
      log.Print(err)
    } else {
      fmt.Printf("Receive data=%#v\n", msg)
    }
  }
  */
}

