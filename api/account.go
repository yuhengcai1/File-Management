package api

import (
	"DB/DB"
	"DB/token"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Username  string		`json:"username"`
	Password  string       `json:"password"`
}

// add admin to the admin list and user list
func (server *Server) createuser(ctx * gin.Context){
	var req CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err!= nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var admin sql.NullBool

	if authPayload.Admin {
    admin = sql.NullBool{Valid: true, Bool: true}
	} else {
    admin = sql.NullBool{Valid: true, Bool: false}
	}


	arg :=  DB.CreateUsersParams{
		Username: req.Username,
        Userhash: req.Password,
		ID:  	authPayload.UserID,
		Admin:  admin,
	}

	user, err := server.store.CreateUsers(ctx,arg); 

	if err!= nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

	rep := NewUserResponse(user)

	ctx.JSON(http.StatusOK, rep)

}



type GetUserByIDAndAdminParams struct {
	ID    int32        `uri:"id"`
}

func (server *Server) getuser(ctx *gin.Context){
	var req GetUserByIDAndAdminParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
        return
	}
	
	user1, err := server.store.GetUserByIDAndAdmin(ctx, req.ID)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.Admin == false{
		err := errors.New("Not admin")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.Admin == false{
		err := errors.New("Not admin")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.Admin == false{
		err := errors.New("Not admin")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.Admin == true{
		ctx.JSON(http.StatusOK, document)
		return
	}else if authPayload.UserID == document.Createdby.Int32{
		ctx.JSON(http.StatusOK, document)
		return
	}else{
		err := errors.New("Not admin")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.Admin == true{
		ctx.JSON(http.StatusOK, document)
		return
	}else if authPayload.UserID == document.Createdby.Int32{
		ctx.JSON(http.StatusOK, document)
		return
	}else{
		err := errors.New("not admin")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
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

	document, err := server.store.GetDocumentByID(ctx, req.documentid)

	if err != nil {
		ctx.JSON(400, gin.H{"error": errorResponse(err)})
		return
	}
	
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.Admin == true{
		err := server.store.DeleteDocumentAdmin(ctx, req.documentid)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}else if authPayload.UserID == document.Createdby.Int32{
		err := server.store.DeleteDocumentAdmin(ctx, req.documentid)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
	}else{
		err := errors.New("not admin")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	

}



