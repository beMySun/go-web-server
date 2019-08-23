package handler

import (
	"go-web-server/dao"
)

type Registry struct {
	DAO *dao.DAO
}

func NewRegistry() *Registry {
	r := &Registry{
		DAO: dao.NewDAO(),
	}
	return r
}
