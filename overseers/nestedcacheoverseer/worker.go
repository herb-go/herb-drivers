package nestedcacheoverseer

import (
	"github.com/herb-go/datamodules/ncache"
	"github.com/herb-go/worker"
)

var cacheworker = ncache.NewNestedCache()
var Team = worker.GetWorkerTeam(&cacheworker)

func GetNestedCacheByID(id string) *ncache.NestedCache {
	w := worker.FindWorker(id)
	if w == nil {
		return nil
	}
	c, ok := w.Interface.(**ncache.NestedCache)
	if ok == false || c == nil {
		return nil
	}
	return *c
}
