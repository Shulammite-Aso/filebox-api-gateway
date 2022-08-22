package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Shulammite-Aso/filebox-api-gateway/pkg/auth/proto"
	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &proto.ValidateRequest{
		Token: token[1],
	})

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if res.Error != "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Set("username", res.Username)

	ctx.Next()
}

func (c *AuthMiddlewareConfig) VerifyReceiver(ctx *gin.Context) {

	// Get reciever username from the request, then use it to request for email of the user from auth service

	type SendFileToPersonRequestBody struct {
		FileName         string `json:"fileName"`
		File             []byte `json:"file"`
		ReceiverUsername string `json:"receiverUsername"`
	}

	body := SendFileToPersonRequestBody{}

	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))

	err = json.Unmarshal(jsonData, &body)
	if err != nil {
		log.Println("could not unmarshal json data")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res, err := c.svc.Client.VerifyReceiver(context.Background(), &proto.VerifyRequest{
		Username: body.ReceiverUsername,
	})

	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if res.Error != "" {
		log.Println(res.Error)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.Set("email", res.Email)

	ctx.Next()
}
