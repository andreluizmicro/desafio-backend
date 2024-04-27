package gateway

import (
	"encoding/json"
	"fmt"
	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/client"
	"io"
	"net/http"
)

type NotificationGatewayResponse struct {
	Message bool `json:"message"`
}

type NotificationGateway struct {
	notificationClient *client.NotificationClient
}

func NewNotificationGateway() *NotificationGateway {
	return &NotificationGateway{
		notificationClient: client.NewNotificationClient(),
	}
}

func (n *NotificationGateway) Notify() bool {
	url := fmt.Sprintf(
		"%s/%s/%s",
		n.notificationClient.Url, n.notificationClient.Version, n.notificationClient.Token,
	)
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	var response NotificationGatewayResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false
	}

	if !response.Message {
		return false
	}
	return true
}
