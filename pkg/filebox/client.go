package filebox

import (
	"fmt"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/config"
	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/filebox/proto"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client proto.FileboxServiceClient
}

func InitServiceClient(c *config.Config) proto.FileboxServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.FileboxSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return proto.NewFileboxServiceClient(cc)
}
