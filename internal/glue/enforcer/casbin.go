package enforcer

import (
	"github.com/casbin/casbin/v2"
)

type CasbinMiddleware interface {
	GetEnforcer() *casbin.Enforcer
}

type casbinMiddleware struct {
	enforcer *casbin.Enforcer
}

func CasbinInit(enforcer *casbin.Enforcer) CasbinMiddleware {
	return &casbinMiddleware{
		enforcer,
	}
}

func (c casbinMiddleware) GetEnforcer() *casbin.Enforcer {
	return c.enforcer
}
