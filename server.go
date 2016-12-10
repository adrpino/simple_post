package main

// imports with the underscore are just drivers, without it you have full access
import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Info struct {
    Name     string `json:"name"`
    Datetime string `json:"datetime"`
}

type params struct {
    context *gin.Context
    statement *sql.Stmt
    row Info
}

func main() {
    fmt.Println("Starting server")
    db, cn_err := sql.Open("mysql", "root:mysqlroot@(127.0.0.1:3306)/test_db")
    if cn_err != nil {
            panic(cn_err.Error())
    }
    stmtIns, err := db.Prepare("INSERT INTO test (`name`, `time`) VALUES( ?, ? )") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close()

    router := gin.Default()

    router.POST("/post", func(c *gin.Context) {
        var json Info
        if c.BindJSON(&json) == nil {
                par := params{context: c, statement: stmtIns, row: json}
                go InsertDb(par)
        }
    })


    router.Run(":8080")
}

//func InsertDb(stmtIns *sql.Stmt, json Info) {
func InsertDb(par params) {
            rw := par.row
            context := par.context
            stmt:= par.statement
            if len(rw.Datetime) == 0 {
                rw.Datetime = time.Now().Format("2006-01-02 15:04:05")
            }
            _, ins_err := stmt.Exec(rw.Name, rw.Datetime)

            if ins_err != nil {
                fmt.Println(ins_err.Error())
            } else {
                context.Status(http.StatusOK)

            }
}
