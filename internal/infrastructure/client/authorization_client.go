package client

import "github.com/andreluizmicro/desafio-backend/config"

type AuthorizationClient struct {
	Url     string
	Version string
	Token   string
}

func NewAuthorizationClient() *AuthorizationClient {
	authClient := config.GetAuthorizationConfigClient()

	return &AuthorizationClient{
		Url:     authClient.AuthorizationClientApiUrl,
		Version: authClient.AuthorizationClientApiVersion,
		Token:   authClient.AuthorizationClientApiToken,
	}
}
