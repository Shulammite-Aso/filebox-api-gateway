package routes

import (
	"context"
	"net/http"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/filebox/proto"
	"github.com/gin-gonic/gin"
)

type SendFileRequestBody struct {
	File []byte `json:"file"`
}

func SendFile(ctx *gin.Context, c proto.FileboxServiceClient) {
	body := SendFileRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	username, _ := ctx.Get("username")

	res, err := c.SendFile(context.Background(), &proto.SendFileRequest{
		Username: username.(string),
		File:     body.File,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
