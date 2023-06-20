package api

import (
	"DB/DB"
	"DB/token"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var sqlDb *sql.DB

type Server struct {
    store *DB.Store
    router *gin.Engine
    tokerMaker token.Maker
}

func NewServer(store *DB.Store) *Server {
    tokenMaker, err := token.NewJWTMaker("")
    if err != nil {
        return nil
    }
    server := &Server{
        store: store,
        tokerMaker: tokenMaker,
        
    }
    router := gin.Default()

    router.POST("/user", server.createuser)
    router.GET("/user/:id", server.getuser)
    router.DELETE("user",server.deleteuser)
    router.PUT("/user/", server.updateuser)

    router.POST("/documents", server.addDocuments)
    router.GET("/documents/:name", server.getDocumentsByName)
    router.GET("/documents/id", server.getDocumentsByID)
    router.DELETE("/documents/admin/:id", server.deleteDocumentAdmin)
    router.DELETE("/documents/normal/:id", server.deleteDocumentNormal)
   
    return server
}

func Init() {
	var err error
    sqlDb, err = sql.Open("postgres", "postgres://postgres:postgrespw@postgres:5432")
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

type CreateDocumentParams struct {
	Documentid int32         `json:"documentid"`
	Name       string        `json:"name"`
	Createdby  sql.NullInt32 `json:"createdby"`
}




