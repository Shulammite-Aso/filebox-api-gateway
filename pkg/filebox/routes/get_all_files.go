package routes

import (
	"context"
	"net/http"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/filebox/proto"
	"github.com/gin-gonic/gin"
)

func GetListOfAllFiles(ctx *gin.Context, c proto.FileboxServiceClient) {

	username, _ := ctx.Get("username")

	res, err := c.GetListOfAllFiles(context.Background(), &proto.GetListOfAllFilesRequest{
		Username: username.(string),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
