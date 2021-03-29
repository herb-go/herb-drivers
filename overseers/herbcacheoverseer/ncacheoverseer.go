package herbcacheoverseer

import (
	"github.com/herb-go/worker"
)

//Config overseer config struct
type Config struct {
}

//ApplyTo apply config to overseer
func (c *Config) ApplyTo(o *worker.PlainOverseer) error {
	o.WithIntroduction("Ncache workers")
	// o.WithOutsourceFunc(func(o *worker.Oncacheutsourced) error {
	// 	c := ncache.New()
	// 	config := herbcache.NewConfig()
	// 	err := o.TranningPlan(config)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	err = config.ApplyTo(c)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	worker.Hire(o.Name, &c)
	// 	worker.OnStart(func() {
	// 		err := c.Start()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	})
	// 	worker.OnStop(func() {
	// 		err := c.Stop()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	})
	// 	return nil
	// })
	return nil
}

//New create new config
func New() *Config {
	return &Config{}
}
