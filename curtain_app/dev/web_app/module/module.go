package module

import (
  "database/sql"
  "fmt"
  "time"

  _ "github.com/go-sql-driver/mysql"
)

type Log struct {
  Id int          `json: id`
  Date time.Time  `json: date`
  Msg string      `json: msg`
}

func DBConnect() (db *sql.DB) {
    dbDriver := "mysql"
    // dbUser := "curtain_app_user"
    dbUser := "root"
    // dbPass := "curtain_app_PW#0"
    dbPass := "root00"
    dbName := "curtain_app"
    dbProtocol := "tcp(mariadb:3306)"
    dbOption := "?parseTime=true"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbProtocol+"/"+dbName+dbOption)
    if err != nil {
      panic(err.Error())
    }
    return db
}

func DBInsert(msg string) {
  db := DBConnect()
  // date := time.Now().Format("2006-01-02 15:04:05")
  date := time.Now()
  insert_sql := "insert into wake_up_log (date, msg) values (?, ?);"
  _, err := db.Exec(insert_sql, date, msg)
  if err != nil {
    panic(err.Error())
  }
  fmt.Printf("insert: %s(%v), %s(%v).\n", date, date, msg, msg)
}

func DBSelect() []Log {
  db := DBConnect()
  select_sql := "select * from wake_up_log order by id desc"
  res, err := db.Query(select_sql)
  if err != nil {
    panic(err.Error())
  }

  logs := []Log{}
  for res.Next() {
    var id int
    var date time.Time
    var msg string
    res.Scan(&id, &date, &msg)
    logs = append(logs, Log{id, date, msg})
  }
  return logs
}

/*
func main() {
  var str string
  fmt.Scan(&str)

  if str == "insert" {
    DBInsert("hoge")
  }

  logs := DBSelect()
  for i, log := range logs {
    fmt.Printf("log[%d]: %v\n", i, log)
  }
}
*/
