package hiredpersist

import (
	"github.com/herb-go/herb-drivers/overseers/persistoverseer"
	"github.com/herb-go/herb/persist"
	"github.com/herb-go/worker"
)

type Config struct {
	ID     string
	Config func(v interface{}) error `config:", lazyload"`
}

func (c *Config) CreateStore() (persist.Store, error) {
	f := persistoverseer.GetStoreByID(c.ID)
	if f == nil {
		return nil, worker.NewWorkerNotFounderError(c.ID)
	}
	return f(c.Config)
}
