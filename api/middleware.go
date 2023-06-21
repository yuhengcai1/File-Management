package api

import (
	"DB/token"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey = "authorization"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc{
	return func(ctx *gin.Context){
		  authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		  if len(authorizationHeader) == 0 {
			err := errors.New("authortization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		  }
	}
}