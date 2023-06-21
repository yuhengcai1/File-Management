package api

import (
	"DB/DB"
	"DB/util"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type loginUser struct {
	UserID  int32       `json:"username"`
	Password  string       `json:"password"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	User userResponse `json:"user"`
}

type userResponse struct {
	UserID  int32      `json:"username"`
	CreatedAt         time.Time `json:"created_at"`
}

func NewUserResponse(user DB.User) userResponse {
	return userResponse{
		UserID:          user.ID,
		CreatedAt:         user.CreatedAt.Time,
	}
}

func (server *Server) loginUser(ctx *gin.Context){
	var req loginUser 

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	
	user, err := server.store.GetUserByID(ctx,req.UserID)
	if err != nil {
		if err == sql.ErrNoRows { 
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Userhash)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized,errorResponse(err))
		return
	}

	accessToken, _, err := server.tokerMaker.CreateToken(
		user.ID,
		server.config.AccessToken,
		user.Admin.Bool,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}


	rsp := loginResponse{
		AccessToken: accessToken,
		User: NewUserResponse(user),
	}

	ctx.JSON(http.StatusOK,rsp)

}