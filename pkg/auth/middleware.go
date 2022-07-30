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
