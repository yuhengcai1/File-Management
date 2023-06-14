package api

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"net/http"
)

var sqlDb *sql.DB



func init() {
	var err error
    sqlDb, err = sql.Open("postgres", "root:root@tcp(127.0.0.1:3306)/test")
    if err!= nil {
        panic(err)
    }
	r := gin.Default()
	r.GET("/normal", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello",
        })
    })
	r.Run(":0800")
}


