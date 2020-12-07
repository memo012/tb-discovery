package registry

import "github.com/memo012/tb-discovery/model"
type Registry struct {
	gd        *Guard // 设置过期时间
}

func (r *Registry) newApp() *model.App {
	return &model.App{}
}

func (r *Registry) Register(ins *model.Instance, latestTime int64) (err error) {
	a := r.newApp()
	// 向注册中心注册
	ins, ok := a.NewInstance(ins, latestTime)
	// 设置过期时间  自我保护机制
	if ok {
		r.gd.incrExp()
	}
	return
}
