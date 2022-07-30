package filebox

import (
	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/auth"
	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/config"
	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/filebox/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/filebox")
	routes.Use(a.AuthRequired)
	routes.POST("/sendfile", svc.SendFile)
	routes.GET("/getfile/:fileName", svc.GetFile)
	routes.PUT("/updatefile/", svc.UpdateFile)
	routes.DELETE("deletefile/:fileName", svc.DeleteFile)
	routes.POST("/sendfile-to-person", svc.SendFileToPerson)
	routes.GET("/get-list-of-all-files/:username", svc.GetListOfAllFiles)
}

func (svc *ServiceClient) SendFile(ctx *gin.Context) {
	routes.SendFile(ctx, svc.Client)
}

func (svc *ServiceClient) GetFile(ctx *gin.Context) {
	routes.GetFile(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateFile(ctx *gin.Context) {
	routes.UpdateFile(ctx, svc.Client)
}

func (svc *ServiceClient) DeleteFile(ctx *gin.Context) {
	routes.DeleteFile(ctx, svc.Client)
}

func (svc *ServiceClient) SendFileToPerson(ctx *gin.Context) {
	routes.SendFileToPerson(ctx, svc.Client)
}

func (svc *ServiceClient) GetListOfAllFiles(ctx *gin.Context) {
	routes.GetListOfAllFiles(ctx, svc.Client)
}
