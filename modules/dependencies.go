package modules

import (
	_ "github.com/aacfactory/fns-contrib/authorizations/encoding/jwt"
	"github.com/aacfactory/fns/service"
	"github.com/aacfactory/fns/service/builtin/authorizations"
)

func dependencies() (services []service.Service) {
	services = append(
		services,
		authorizations.Service(),
	)
	return
}
