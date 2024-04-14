package network

import (
	"demo-scrapping/authenticator"
	"demo-scrapping/config"
	"demo-scrapping/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	cfg   *config.Config
	engin *gin.Engine

	service       service.ServiceImpl
	authenticator authenticator.AuthenticatorImpl
}

func NewNetwork(
	cfg *config.Config,
	service service.ServiceImpl,
	authenticator authenticator.AuthenticatorImpl,
) *Network {
	n := &Network{
		cfg:           cfg,
		service:       service,
		authenticator: authenticator,
		engin:         gin.New(),
	}

	n.setCors()
	newAdmin(n)

	return n
}

func (n *Network) Run() error {
	return n.engin.Run(n.cfg.Network.Port)
}
