package ncacheoverseer

import (
	"github.com/herb-go/datamodules/ncache"
	"github.com/herb-go/worker"
)

var cacheworker = ncache.New()
var Team = worker.GetWorkerTeam(&cacheworker)

func GetCacheByID(id string) *ncache.Cache {
	w := worker.FindWorker(id)
	if w == nil {
		return nil
	}
	c, ok := w.Interface.(**ncache.Cache)
	if ok == false || c == nil {
		return nil
	}
	return *c
}
