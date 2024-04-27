package gateway

type AuthorizationGateway interface {
	AuthorizeTransfer() bool
}
