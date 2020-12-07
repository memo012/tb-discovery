package discovery

import (
	"github.com/memo012/tb-discovery/model"
)

func (d *Discovery) Register(ins *model.Instance, latestTimestamp int64) {
	// 注册 instance
	_ = d.registry.Register(ins, latestTimestamp)
}

// Fetch fetch all instances by appid.
func (d *Discovery) Fetch(arg *model.ArgFetch) (info *model.InstanceInfo, err error) {
	return d.registry.Fetch(arg.Zone, arg.Env, arg.AppID, 0)
}






