package discovery

import (
	"context"
	"github.com/memo012/tb-discovery/conf"
	"github.com/memo012/tb-discovery/registry"
)

type Discovery struct {
	c        *conf.Config
	registry *registry.Registry
}

// New get a discovery.
func New(c *conf.Config) (d *Discovery, cancel context.CancelFunc) {
	d = &Discovery{
		c:        c,
		registry: registry.NewRegistry(c),
	}
	return
}
