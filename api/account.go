package api

import (
	"DB/DB"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	ID        int32        `json:"id"`
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	Admin   bool 			`json:"admin"`
}

// add admin to the admin list and user list
func (server *Server) createuser(ctx * gin.Context){
	var req CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err!= nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }
	var admin sql.NullBool

	if req.Admin {
    admin = sql.NullBool{Valid: true, Bool: true}
	} else {
    admin = sql.NullBool{Valid: true, Bool: false}
	}


	arg :=  DB.CreateUsersParams{
		Username: req.Username,
        Userhash: req.Password,
		ID:  	req.ID,
		Admin:  admin,
	}

	if _, err := server.store.CreateUsers(ctx,arg); err!= nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

}



type GetUserByIDAndAdminParams struct {
	ID    int32        `uri:"id"`
	Admin sql.NullBool `uri:"admin"`
}

func (server *Server) getuser(ctx *gin.Context){
	var req GetUserByIDAndAdminParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
        return
	}

	arg := DB.GetUserByIDAndAdminParams{
		ID: req.ID,
		Admin: req.Admin,
	}
	
	user1, err := server.store.GetUserByIDAndAdmin(ctx, arg)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user1)

}


type UpdateUsersParams struct {
	Userhash string `json:"userhash"`
	ID       int32  `json:"id"`
}

func (server *Server) updateuser(ctx *gin.Context){
	var req UpdateUsersParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
        return
	}

	arg := DB.UpdateUsersParams{
		Userhash: req.Userhash,
		ID: req.ID,
	}
	
	err := server.store.UpdateUsers(ctx, arg)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	
}


type Deleteuser struct {
	ID int32 `json:"id" binding:"required,min=1"`
}


func (server *Server) deleteuser(ctx *gin.Context){
	var req Deleteuser

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
        return
	}

	err := server.store.Deleteusers(ctx, int32(req.ID))

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	
}

type ADDDocumentParams struct {
	Documentid int32         `json:"documentid"`
	Name       string        `json:"name"`
	Createdby  sql.NullInt32 `json:"createdby"`
}

func (server *Server) addDocuments(ctx *gin.Context){
	var req ADDDocumentParams

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
        return
	}
	arg:= DB.CreateDocumentParams{
		Documentid: req.Documentid,
		Name: req.Name,
		Createdby : req.Createdby,
	}
	
	document, err := server.store.CreateDocument(ctx, arg)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusOK, document)
}



type GetDocumentById struct {
	ID int32 `JSON:"id" binding:"required,min=0"`
}

func (server *Server) getDocumentsByID(ctx *gin.Context){
	var req GetDocumentById

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
        return
	}

	document, err := server.store.GetDocumentByID(ctx, req.ID)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusOK, document)
}



type GetDocumentByName struct {
	Name       string        `json:"name"`
}

func (server *Server) getDocumentsByName(ctx *gin.Context){
	var req GetDocumentById

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
        return
	}

	document, err := server.store.GetDocumentByID(ctx, req.ID)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}


	ctx.JSON(http.StatusOK, document)
}


type deleteDocumentAdmin struct {
	documentid int32 `json:"documentid"`
}

func (server *Server) deleteDocumentAdmin(ctx *gin.Context){
	var req deleteDocumentAdmin

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
        return
	}

	err := server.store.DeleteDocumentAdmin(ctx, req.documentid)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

}


type deleteDocumentNormal struct {
	documentid int32 `json:"documentid"`
	Createdby  sql.NullInt32 `json:"createdby"`
}

func (server *Server) deleteDocumentNormal(ctx *gin.Context){
	var req deleteDocumentNormal

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
        return
	}

	arg := DB.DeleteDocumentNormalParams{
		Documentid: req.documentid,
		Createdby: req.Createdby,
	}
	
	err := server.store.DeleteDocumentNormal(ctx, arg)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

}


