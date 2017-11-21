package cli

import "goconfig"

type Cfg struct {
	Section string
	c       *goconfig.ConfigFile
}

func NewCfg(fileName string) (*Cfg, error) {
	cfg, err := goconfig.LoadConfigFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Cfg{c: cfg}, nil
}

func (cfg *Cfg) GetString(key string) (string, error) {
	return cfg.c.GetValue(cfg.Section, key)
}

func (cfg *Cfg) GetInt(key string) (int64, error) {
	return cfg.c.Int64(cfg.Section, key)
}

func (cfg *Cfg) GetFloat(key string) (float64, error) {
	return cfg.c.Float64(cfg.Section, key)
}
