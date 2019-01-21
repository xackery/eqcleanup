package client

import (
	"github.com/pkg/errors"
	"github.com/xackery/eqcleanup/eqdb"
	"github.com/xackery/eqcleanup/quest"
	"github.com/xackery/eqemuconfig"
)

// Client represents eqcleanup
type Client struct {
	Database *eqdb.Database
	Quest    *quest.Quest
	Config   *eqemuconfig.Config
}

// New creates a new instance of Client
func New() (c *Client, err error) {
	c = &Client{}

	c.Config, err = eqemuconfig.GetConfig()
	if err != nil {
		err = errors.Wrap(err, "failed to get eqemu config")
		return
	}
	c.Database, err = eqdb.New(c.Config)
	if err != nil {
		err = errors.Wrap(err, "failed to connect to database")
		return
	}
	c.Quest, err = quest.New(c.Config)
	if err != nil {
		err = errors.Wrap(err, "failed to find quest directory")
		return
	}
	return
}
