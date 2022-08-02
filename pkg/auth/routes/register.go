package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/auth/proto"
	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c proto.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.Register(context.Background(), &proto.RegisterRequest{
		Email:    body.Email,
		Username: body.Username,
		Password: body.Password,
	})

	if err != nil {
		fmt.Println("Error msg:", err)
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	fmt.Println("Register response:", res.Message)

	ctx.JSON(http.StatusCreated, &res)
}
