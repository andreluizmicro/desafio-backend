package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/client"
)

const (
	operationAuthorized = "Autorizado"
)

type AuthorizationGatewayResponse struct {
	Message string `json:"message"`
}

type AuthorizationGateway struct {
	authorizationClient *client.AuthorizationClient
}

func NewAuthorizationGateway() *AuthorizationGateway {
	return &AuthorizationGateway{
		authorizationClient: client.NewAuthorizationClient(),
	}
}

func (g *AuthorizationGateway) AuthorizeTransfer() bool {
	url := fmt.Sprintf(
		"%s/%s/%s",
		g.authorizationClient.Url, g.authorizationClient.Version, g.authorizationClient.Token,
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

	var response AuthorizationGatewayResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false
	}
	if response.Message != operationAuthorized {
		return false
	}
	return true
}
