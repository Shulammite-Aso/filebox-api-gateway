package routes

import (
	"context"
	"net/http"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/filebox/proto"
	"github.com/gin-gonic/gin"
)

type SendFileToPersonRequestBody struct {
	File     []byte `json:"file"`
	Username string `json:"username"`
}

func SendFileToPerson(ctx *gin.Context, c proto.FileboxServiceClient) {
	body := SendFileToPersonRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	username, _ := ctx.Get("username")

	res, err := c.SendFileToPerson(context.Background(), &proto.SendFileToPersonRequest{
		Username:       body.Username,
		File:           body.File,
		SenderUsername: username.(string),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
