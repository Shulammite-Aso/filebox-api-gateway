package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/auth/proto"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &proto.ValidateRequest{
		Token: token[1],
	})

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("username", res.Username)

	ctx.Next()
}

func (c *AuthMiddlewareConfig) VerifyReceiver(ctx *gin.Context) {

	// Get reciever username from the request, then use it to request for email of the user from auth service

	receiverUsername, _ := ctx.Get("receiverUsername")

	res, err := c.svc.Client.VerifyReceiver(context.Background(), &proto.VerifyRequest{
		Username: receiverUsername.(string),
	})

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.Set("email", res.Email)

	ctx.Next()
}
