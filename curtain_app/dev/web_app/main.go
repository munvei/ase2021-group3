package main

import (
  "net/http"
  "log"

  "github.com/gin-gonic/gin"
  "gopkg.in/olahol/melody.v1"

  "./module"
)

func main() {
  r := gin.Default()
  // 静的ファイルの設定
  path := "/go/web_app"
  r.LoadHTMLGlob(path+"/templates/*")
  r.Static("/js", path+"/js")
  m := melody.New()

  // User Agentを取得
  var ua string
  r.Use(func(c *gin.Context) {
    ua = c.GetHeader("User-Agent")
    c.Next()
  })

  // ルーティング
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "User-Agent": ua,
      "message": "Hello World",
    })
  })

  // react test
  r.GET("/react", func(c *gin.Context) {
    c.HTML(http.StatusOK, "react-test.html", gin.H{})
  })

  // selectした内容を表示
  r.GET("/db", func(c *gin.Context) {
    c.HTML(http.StatusOK, "showdb.html", gin.H{"data": module.DBSelect()})
  })

  // selectした内容をjsonで返す
  r.GET("/db_json", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "rows": module.DBSelect(),
    })
  })

  // dbに追加
  r.POST("/insert", func(c *gin.Context) {
    msg := c.PostForm("msg")
    module.DBInsert(msg)
    c.JSON(http.StatusOK, gin.H{
      "status": "OK",
      "msg": msg,
    })
  })

  // LINEにメッセージを送る
  r.POST("/line", func(c *gin.Context) {
    msg := c.PostForm("msg")
    res := module.SendLine(path+"/module/token.txt", msg)
    c.JSON(http.StatusOK, gin.H{
      "res": res,
    })
  })

  // websocketのテスト
  r.GET("/test", func(c *gin.Context) {
    c.HTML(http.StatusOK, "test.html", gin.H{
      "msg": "websocket test",
    })
  })

  r.GET("/test/ws", func(c *gin.Context) {
    m.HandleRequest(c.Writer, c.Request)
  })
  m.HandleMessage(func(s *melody.Session, msg []byte) {
    m.Broadcast(msg)
  })

  m.HandleConnect(func(s *melody.Session) {
    log.Printf("websocket connection open. [session: %#v]\n", s)
  })
  m.HandleDisconnect(func(s *melody.Session) {
    log.Printf("websocket connection close. [session: %#v]\n", s)
  })

  r.Run() // listen and serve on 0.0.0.0:8080 <- default
}

/*
// 動かん
func handleWebSocket(c *gin.Context) {
  c.HTML(http.StatusOK, "ws.html", gin.H{
    "msg": "websocket test",
  })

  websocket.Handler(func(ws *websocket.Conn) {
    defer ws.Close()

    // first message
    msg := "How are you?"
    err := websocket.Message.Send(ws, msg)
    if err != nil {
      panic(err.Error())
    }

    for {
      var msg string
      // receive client message
      err = websocket.Message.Receive(ws, &msg)
      if err != nil {
        panic(err.Error())
      }

      // send message
      err = websocket.Message.Send(ws, fmt.Sprintf("Received msg: \"%s\"", msg))
      if err != nil {
        panic(err.Error())
      }
    }
  }).ServeHTTP(c.Writer, c.Request)
}
*/
