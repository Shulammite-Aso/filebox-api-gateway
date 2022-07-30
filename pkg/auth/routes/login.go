package routes

import (
	"context"
	"net/http"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/auth/proto"
	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c proto.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &proto.LoginRequest{
		Username: b.Username,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
