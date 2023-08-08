package client

import (
	"log"
	"tclient/internal/config"

	"github.com/anonyindian/gotgproto"
	"github.com/anonyindian/gotgproto/sessionMaker"
)

func Create(c *config.AppConfig) (*gotgproto.Client, error) {
	clientType := gotgproto.ClientType{
		Phone: c.Phone,
	}
	client, err := gotgproto.NewClient(
		c.AppId,
		c.ApiHash,
		clientType,
		&gotgproto.ClientOpts{
			Session:          sessionMaker.NewSession("tclient", sessionMaker.Session),
			DisableCopyright: true,
		},
	)
	if err != nil {
		log.Fatalln("failed to start client:", err)
	}

	return client, err
}
