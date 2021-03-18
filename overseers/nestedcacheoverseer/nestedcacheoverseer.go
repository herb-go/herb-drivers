package nestedcacheoverseer

import (
	"github.com/herb-go/datamodules/ncache"
	"github.com/herb-go/datamodules/ncache/builderconfig"
	"github.com/herb-go/worker"
)

//Config overseer config struct
type Config struct {
}

func train(c *ncache.NestedCache, loader func(v interface{}) error) error {
	config := &builderconfig.BuildConfig{}
	err := loader(config)
	builders, err := config.CreateBuilders()
	if err != nil {
		return err
	}
	c.WithBuilder(builders...)
	return nil
}

//ApplyTo apply config to overseer
func (c *Config) ApplyTo(o *worker.PlainOverseer) error {
	o.WithIntroduction("NestedCache workers")
	o.WithTrainFunc(func(w []*worker.Worker) error {
		for _, v := range w {
			cache := GetNestedCacheByID(v.Name)
			if cache == nil {
				continue
			}
			t := worker.GetTranning(v.Name)
			if t == nil {
				continue
			}
			err := train(cache, t.TranningPlan)
			if err != nil {
				return err
			}
			return nil
		}
		return nil
	})
	o.WithOutsourceFunc(func(o *worker.Outsourced) error {
		c := ncache.NewNestedCache()
		err := train(c, o.TranningPlan)
		if err != nil {
			return err
		}
		worker.Hire(o.Name, &c)
		worker.OnStart(func() {
			err := c.Start()
			if err != nil {
				panic(err)
			}
		})
		worker.OnStop(func() {
			err := c.Stop()
			if err != nil {
				panic(err)
			}
		})
		return nil
	})
	return nil
}

//New create new config
func New() *Config {
	return &Config{}
}
