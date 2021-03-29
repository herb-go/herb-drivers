package herbcacheoverseer

import (
	"github.com/herb-go/datamodules/herbcache"
	"github.com/herb-go/worker"
)

var cacheworker = herbcache.New()
var Team = worker.GetWorkerTeam(&cacheworker)

func GetCacheByID(id string) *herbcache.Cache {
	w := worker.FindWorker(id)
	if w == nil {
		return nil
	}
	c, ok := w.Interface.(**herbcache.Cache)
	if ok == false || c == nil {
		return nil
	}
	return *c
}
