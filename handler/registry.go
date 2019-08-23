package handler

import (
	"go-web-server/dao"
)

// Registry defination
type Registry struct {
	DAO *dao.DAO
}

// NewRegistry defination
func NewRegistry() *Registry {
	r := &Registry{
		DAO: dao.NewDAO(),
	}
	return r
}
