package eqconfig

import (
	"github.com/xackery/eqemuconfig"
)

var config *eqemuconfig.Config

func Load() (cfg *eqemuconfig.Config, err error) {
	if config != nil {
		cfg = config
		return
	}
	if config, err = eqemuconfig.GetConfig(); err != nil {
		return
	}
	cfg = config
	return
}
