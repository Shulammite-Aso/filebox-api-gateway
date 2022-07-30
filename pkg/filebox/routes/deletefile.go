package routes

import (
	"context"
	"net/http"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/filebox/proto"
	"github.com/gin-gonic/gin"
)

func DeleteFile(ctx *gin.Context, c proto.FileboxServiceClient) {

	username, _ := ctx.Get("username")
	fileName := ctx.Param("fileName")

	res, err := c.DeleteFile(context.Background(), &proto.DeleteFileRequest{
		Username: username.(string),
		FileName: fileName,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
