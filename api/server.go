package api

import (
	"DB/DB"
	"DB/token"
	"DB/util"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var sqlDb *sql.DB

type Server struct {
    config util.Config
    store *DB.Store
    router *gin.Engine
    tokerMaker token.Maker
}

func NewServer(config util.Config, store *DB.Store) (*Server,error) {
    tokenMaker, err := token.NewJWTMaker(config.TokenSymmetriKEy) //

    if err != nil {
        return nil,fmt.Errorf("cannot create token %w", err)
    }

    server := &Server{
        config : config,
        store: store,
        tokerMaker: tokenMaker,
    }

    router := gin.Default()

    authRoutes := router.Group("/").Use(authMiddleware(server.tokerMaker))

    authRoutes.POST("/user", server.createuser)
    authRoutes.GET("/user/:id", server.getuser)
    authRoutes.DELETE("user",server.deleteuser)
    authRoutes.PUT("/user/", server.updateuser)

    router.POST("/user/login", server.loginUser)

    authRoutes.POST("/documents", server.addDocuments)
    authRoutes.GET("/documents/:name", server.getDocumentsByName)
    authRoutes.GET("/documents/id", server.getDocumentsByID)
    authRoutes.DELETE("/documents/:id", server.deleteDocumentAdmin)
   

    server.router = router

    return server, nil
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

func (server *Server) Start(address string) error {
    if(server.router == nil) {
        fmt.Println("null in here")
    }
    return server.router.Run(address)
}


func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

