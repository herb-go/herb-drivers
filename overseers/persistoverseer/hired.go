package persistoverseer

import (
	"github.com/herb-go/herb/persist"
	"github.com/herb-go/worker"
)

var persistfactoryworker func(loader func(v interface{}) error) (persist.Store, error)

var Team = worker.GetWorkerTeam(&persistfactoryworker)

func GetStoreByID(id string) func(loader func(v interface{}) error) (persist.Store, error) {
	w := worker.FindWorker(id)
	if w == nil {
		return nil
	}
	c, ok := w.Interface.(*func(loader func(v interface{}) error) (persist.Store, error))
	if ok == false || c == nil {
		return nil
	}
	return *c
}
