package client

import "github.com/andreluizmicro/desafio-backend/config"

type NotificationClient struct {
	Url     string
	Version string
	Token   string
}

func NewNotificationClient() *NotificationClient {
	authClient := config.GetAuthorizationConfigClient()

	return &NotificationClient{
		Url:     authClient.NotificationClientApiUrl,
		Version: authClient.NotificationClientApiVersion,
		Token:   authClient.NotificationClientApiToken,
	}
}
