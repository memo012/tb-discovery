package registry

import (
	"fmt"
	"github.com/memo012/tb-discovery/constant/errors"
	"github.com/memo012/tb-discovery/model"
	"sync"
)

type Registry struct {
	appm  map[string]*model.Apps // appid-env -> apps    获取instance  那一台机器
	gd    *Guard                 // 设置过期时间
	aLock sync.RWMutex
	a     model.App
}

func (r *Registry) newApp(ins *model.Instance) (a *model.App) {
	as, _ := r.newapps(ins.AppID, ins.Env)
	a, _ = as.NewApp(ins.Zone, ins.AppID, ins.LatestTimestamp)
	return
}

func (r *Registry) newapps(appid, env string) (a *model.Apps, ok bool) {
	key := appsKey(appid, env)
	r.aLock.Lock()
	if a, ok = r.appm[key]; !ok {
		a = model.NewApps()
		r.appm[key] = a
	}
	r.aLock.RUnlock()
	return
}

func appsKey(appid, env string) string {
	return fmt.Sprintf("%s-%s", appid, env)
}

func (r *Registry) Register(ins *model.Instance, latestTime int64) (err error) {
	a := r.newApp(ins)
	// 向注册中心注册

	ins, ok := a.NewInstance(ins, latestTime)
	// 设置过期时间  自我保护机制
	if ok {
		r.gd.incrExp()
	}
	return
}

func (r *Registry) Fetch(zone, env, appid string, latestTime int64) (info *model.InstanceInfo, err error) {
	// appID+env
	key := appsKey(appid, env)
	r.aLock.RLock()
	a, ok := r.appm[key]
	r.aLock.RUnlock()
	if !ok {
		err = errors.NothingFound
		return
	}
	info, err = a.InstanceInfo(zone, latestTime)
	if err != nil {
		return
	}
	return
}
