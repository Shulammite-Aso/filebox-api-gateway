package routes

import (
	"context"
	"net/http"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/filebox/proto"
	"github.com/gin-gonic/gin"
)

type UpdateFileRequestBody struct {
	FileName string `json:"fileName"`
	File     []byte `json:"file"`
}

func UpdateFile(ctx *gin.Context, c proto.FileboxServiceClient) {
	body := UpdateFileRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	username, _ := ctx.Get("username")
	res, err := c.UpdateFile(context.Background(), &proto.UpdateFileRequest{
		Username: username.(string),
		FileName: body.FileName,
		File:     body.File,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
