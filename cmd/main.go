package main

import (
	"log"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/auth"
	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/config"
	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/filebox"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	filebox.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
