package nestedcacheoverseer

import (
	"github.com/herb-go/datamodules/ncache/builderconfig"
	"github.com/herb-go/worker"
)

//Config overseer config struct
type Config struct {
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
			config := &builderconfig.BuildConfig{}
			err := t.TranningPlan(config)
			if err != nil {
				return err
			}
			builders, err := config.CreateBuilders()
			if err != nil {
				return err
			}
			cache.WithBuilder(builders...)
		}
		return nil
	})
	return nil
}

//New create new config
func New() *Config {
	return &Config{}
}
