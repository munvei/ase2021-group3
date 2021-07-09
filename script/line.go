package main

import (
  "strings"
  "bytes"
  "encoding/json"
  "fmt"
  "net/http"
  "io/ioutil"
)

type Msg struct {
  Type string `json: "type"`
  Text string     `json: "text"`
}

type Msgs struct {
  Messages []Msg      `json: "messages"`
}

func main() {
  var text string
  fmt.Scan(&text)
  msg := Msg{"text", text}
  msgs := Msgs{[]Msg{msg}}
  SendMsg(msgs)
}

func SendMsg(msgs Msgs) {
  // read token
  token_path := "./token.txt"
  token_b, read_err := ioutil.ReadFile(token_path)
  if read_err != nil {
    panic(read_err)
  }
  msgs_json_byte, _ := json.Marshal(msgs)
  msgs_json := string(msgs_json_byte)
  msgs_json = strings.ToLower(msgs_json)
  token := strings.TrimRight(string(token_b), "\n")

  URL := "https://api.line.me/v2/bot/message/broadcast"
  req, _ := http.NewRequest("POST", URL, bytes.NewBuffer([]byte(msgs_json)))
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer "+token)

  client := new(http.Client)
  res, post_err := client.Do(req)
  if post_err != nil {
    panic(post_err)
  }
  fmt.Println(bytes.NewBuffer([]byte(msgs_json)))
  fmt.Printf("res: %v\n", res)
}
