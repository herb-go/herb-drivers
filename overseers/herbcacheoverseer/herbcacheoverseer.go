package herbcacheoverseer

import (
	"github.com/herb-go/datamodule-drivers/storageconfig"
	"github.com/herb-go/datamodules/herbcache"
	"github.com/herb-go/worker"
)

type CacheConfig struct {
	Storage   storageconfig.Directive
	Namespace *string
	Group     *string
}

func (c *CacheConfig) ApplyTo(cache *herbcache.Cache) error {
	s := herbcache.NewStorage()
	err := c.Storage.ApplyTo(s)
	herbcache.SetCacheStorage(cache, s)
	if err != nil {
		return err
	}
	if c.Namespace != nil {
		herbcache.SetCacheNamespace(cache, []byte(*c.Namespace))
	}
	if c.Group != nil {
		herbcache.SetCacheGroup(cache, []byte(*c.Group))
	}
	return nil
}

//Config overseer config struct
type Config struct {
}

//ApplyTo apply config to overseer
func (c *Config) ApplyTo(o *worker.PlainOverseer) error {
	o.WithIntroduction("Herbcache workers")
	o.WithOutsourceFunc(func(o *worker.Outsourced) error {
		c := herbcache.New()
		config := &CacheConfig{}
		err := o.TranningPlan(config)
		if err != nil {
			return err
		}
		err = config.ApplyTo(c)
		if err != nil {
			return err
		}
		worker.Hire(o.Name, &c)
		s := c.Storage()
		worker.OnStart(func() {
			err := s.Start()
			if err != nil {
				panic(err)
			}
		})
		worker.OnStop(func() {
			err := s.Stop()
			if err != nil {
				panic(err)
			}
		})
		return nil
	})
	o.WithTrainFunc(func(w []*worker.Worker) error {
		for _, v := range w {
			c := GetCacheByID(v.Name)
			if c == nil {
				continue
			}
			t := worker.GetTranning(v.Name)
			if t == nil {
				continue
			}
			config := &CacheConfig{}
			err := t.TranningPlan(config)
			if err != nil {
				return err
			}
			err = config.ApplyTo(c)
			if err != nil {
				return err
			}
			s := c.Storage()
			worker.OnStart(func() {
				err := s.Start()
				if err != nil {
					panic(err)
				}
			})
			worker.OnStop(func() {
				err := s.Stop()
				if err != nil {
					panic(err)
				}
			})
		}
		return nil
	})
	return nil
}

//New create new config
func New() *Config {
	return &Config{}
}
