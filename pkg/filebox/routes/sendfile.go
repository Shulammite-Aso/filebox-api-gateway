package routes

import (
	"context"
	"net/http"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/filebox/proto"
	"github.com/gin-gonic/gin"
)

type SendFileRequestBody struct {
	Username string `json:"username"`
	File     []byte `json:"file"`
}

func SendFile(ctx *gin.Context, c proto.FileboxServiceClient) {
	body := SendFileRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.SendFile(context.Background(), &proto.SendFileRequest{
		Username: body.Username,
		File:     body.File,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
