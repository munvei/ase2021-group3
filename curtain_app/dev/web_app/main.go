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

  // ルーティング
  r.GET("/", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index-react.html", gin.H{})
  })

  // selectした内容をjsonで返す
  r.GET("/db", func(c *gin.Context) {
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

