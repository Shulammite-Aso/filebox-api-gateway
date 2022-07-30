package auth

import (
	"fmt"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/auth/proto"
	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client proto.AuthServiceClient
}

func InitServiceClient(c *config.Config) proto.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return proto.NewAuthServiceClient(cc)
}
