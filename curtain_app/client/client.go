package main

import (
  "log"
  "fmt"
  "time"
  "regexp"
  "os/exec"

  "golang.org/x/net/websocket"
)

var (
  locate = "localhost/test"
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

  time.Sleep(1000*time.Second)
  sendMsg(ws, "Bye.")
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
    //var keyph=`call`
    r := regexp.MustCompile(`call:`)
    if(r.MatchString(msg)==true){
      msg = r.ReplaceAllString(msg, "./")
      callScript(msg)
    } else {
    fmt.Printf("Receive data=%#v\n", msg)
    }
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

func callScript(msg string){
  err := exec.Command("sh","-c",msg).Run()
  if err != nil {
    fmt.Printf("%#vなどというファイルはないです\n", msg)
  }
}
