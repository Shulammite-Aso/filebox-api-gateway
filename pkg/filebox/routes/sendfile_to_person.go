package routes

import (
	"context"
	"net/http"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/filebox/proto"
	"github.com/gin-gonic/gin"
)

type SendFileToPersonRequestBody struct {
	FileName         string `json:"fileName"`
	File             []byte `json:"file"`
	ReceiverUsername string `json:"receiverUsername"`
}

func SendFileToPerson(ctx *gin.Context, c proto.FileboxServiceClient) {
	body := SendFileToPersonRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	username, _ := ctx.Get("username")
	email, _ := ctx.Get("email")

	res, err := c.SendFileToPerson(context.Background(), &proto.SendFileToPersonRequest{
		Username:       body.ReceiverUsername,
		FileName:       body.FileName,
		Email:          email.(string),
		File:           body.File,
		SenderUsername: username.(string),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
