package routes

import (
	"context"
	"net/http"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/filebox/proto"
	"github.com/gin-gonic/gin"
)

type GetFileRequestBody struct {
	Username string `json:"username"`
	FileName string `json:"file"`
}

func GetFile(ctx *gin.Context, c proto.FileboxServiceClient) {
	body := GetFileRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.GetFile(context.Background(), &proto.GetFileRequest{
		Username: body.Username,
		FileName: body.FileName,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
