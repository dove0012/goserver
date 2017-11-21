package utils

import (
	"goconfig"
)

type Cfg struct {
	Section string
	c       *goconfig.ConfigFile
}

func NewCfg(fileName string) (*Cfg) {
	cfg, err := goconfig.LoadConfigFile("config/" + fileName)
	Log.FailOnError(err, "goconfig.LoadConfigFile error")
	return &Cfg{c: cfg}
}

func (cfg *Cfg) GetString(key string) (string) {
	v, err := cfg.c.GetValue(cfg.Section, key)
	Log.FailOnError(err, "goconfig.GetValue error")
	return v
}

func (cfg *Cfg) GetInt(key string) (int64) {
	v, err := cfg.c.Int64(cfg.Section, key)
	Log.FailOnError(err, "goconfig.Int64 error")
	return v
}

func (cfg *Cfg) GetFloat(key string) (float64) {
	v, err := cfg.c.Float64(cfg.Section, key)
	Log.FailOnError(err, "goconfig.Float64 error")
	return v
}
