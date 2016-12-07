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
    Time string `json:"time"`
}


func main() {
	db, cn_err := sql.Open("mysql", "root:mysqlroot@(127.0.0.1:3306)/test_db")
	if cn_err != nil {
		panic(cn_err.Error())
	}
    stmtIns, err := db.Prepare("INSERT INTO test (`name`, `time`) VALUES( ?, ? )") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

    router := gin.Default()

    router.POST("/post", func(c *gin.Context) {
        var json Info
//		t0 := time.Now()
//		fmt.Println(t0)
		// This basically means that the json could be parsed
        if c.BindJSON(&json) == nil {
			fmt.Println(json.Name, json.Time)
			c.JSON(http.StatusOK, gin.H{"status": "all good"})
//			_, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
        }
    })


    router.Run(":8080")
}
