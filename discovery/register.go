package discovery

import (
	"github.com/memo012/tb-discovery/model"
)

func (d *Discovery) Register(ins *model.Instance, latestTimestamp int64) {
	// 注册 instance
	_ = d.registry.Register(ins, latestTimestamp)
}


