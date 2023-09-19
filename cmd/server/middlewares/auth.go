package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	publicKey, privateKey string
}

func NewAuth(publicKey, privateKey string) *Auth {
	return &Auth{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (a *Auth) AuthHeader(ctx *gin.Context) {
	headerPublicKey := ctx.GetHeader("PUBLIC-KEY")
	privateKey := ctx.GetHeader("PRIVATE-KEY")

	if a.publicKey != headerPublicKey || a.privateKey != privateKey {
		ctx.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized access"))
	}
}
