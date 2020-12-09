package model

import (
	"encoding/json"
	log "github.com/go-kratos/kratos/pkg/log"
	"github.com/memo012/tb-discovery/constant/errors"
	"sync"
	"time"
)

type Instance struct {
	Env      string            `json:"env"`
	Zone     string            `json:"zone"`
	HostName string            `json:"hostname"`
	Metadata map[string]string `json:"metadata"`
	AppID    string            `json:"appid"` // 服务名标识

	// timestamp
	RegTimestamp int64 `json:"reg_timestamp"`
	UpTimestamp  int64 `json:"up_timestamp"`

	LatestTimestamp int64 `json:"latest_timestamp"` // 服务最新更新时间
}

type Apps struct {
	apps            map[string]*App
	lock            sync.RWMutex
	latestTimestamp int64
}

// NewApps return new Apps.
func NewApps() *Apps {
	return &Apps{
		apps: make(map[string]*App),
	}
}

// NewInstance new a instance.
func NewInstance(arg *ArgRegister) (i *Instance) {
	now := time.Now().UnixNano()
	i = &Instance{
		Zone:            arg.Zone,
		Env:             arg.Env,
		AppID:           arg.AppID,
		HostName:        arg.Hostname,
		RegTimestamp:    now,
		UpTimestamp:     now,
		LatestTimestamp: now,
	}
	if arg.Metadata != "" {
		if err := json.Unmarshal([]byte(arg.Metadata), &i.Metadata); err != nil {
			log.Error("json unmarshal metadata err %v", err)
		}
	}
	return
}

type App struct {
	AppID           string
	Zone            string
	lock            sync.RWMutex
	instances       map[string]*Instance // instance名称  key: 主机名 value：instance
	latestTimestamp int64                // 最新更新时间
}

// NewApp new App.
func NewApp(zone, appid string) (a *App) {
	a = &App{
		AppID:     appid,
		Zone:      zone,
		instances: make(map[string]*Instance),
	}
	return
}

// InstanceInfo the info get by consumer.
// 消费者获取的信息
type InstanceInfo struct {
	Instances       map[string][]*Instance `json:"instances"`
	LatestTimestamp int64                  `json:"latest_timestamp"`
}

// NewInstance new a instance.
func (a *App) NewInstance(ins *Instance, latestTime int64) (i *Instance, ok bool) {
	i = new(Instance)
	a.lock.Lock()

	_, ok = a.instances[ins.HostName]
	if ok {
		// instance 存在
		ins.UpTimestamp = ins.RegTimestamp
	}
	a.instances[ins.HostName] = ins
	// 更新时间戳
	a.updateLatest(latestTime)
	*i = *ins
	a.lock.Unlock()
	ok = !ok
	return
}

func (p *Apps) NewApp(zone, appid string, lts int64) (a *App, new bool) {
	p.lock.Lock()
	a, ok := p.apps[zone]
	if !ok {
		a = NewApp(zone, appid)
		p.apps[zone] = a
	}
	if lts <= p.latestTimestamp {
		// insure increase
		lts = p.latestTimestamp + 1
	}
	a.latestTimestamp = lts
	p.lock.Unlock()
	new = !ok
	return
}

func (a *App) updateLatest(latestTime int64) {
	if latestTime <= a.latestTimestamp {
		a.latestTimestamp = a.latestTimestamp + 1
	}
	a.latestTimestamp = latestTime
}

func (p *Apps) InstanceInfo(zone string, latestTime int64) (ci *InstanceInfo, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	// 看instance 是否被改动
	//if latestTime >= p.latestTimestamp {
	//	return nil, errors.NotModified
	//}

	// 说明instance 集合被改动过 拉取修改的instance  定义返回值

	ci = &InstanceInfo{
		LatestTimestamp: p.latestTimestamp,
		Instances:       make(map[string][]*Instance),
	}
	var ok bool
	for z, app := range p.apps {
		if zone == "" || z == zone {
			ok = true
			instances := make([]*Instance, 0)
			for _, i := range app.instances {
				// todo 判断某个instance是否被修改 isModify()
				v := i
				instances = append(instances, v)
			}
			ci.Instances[z] = instances
		}
	}

	if !ok {
		err = errors.NothingFound
	} else if len(ci.Instances) == 0 {
		err = errors.NotModified
	}
	return
}

func isModify(l1, l2 int64) bool {
	return l1 >= l2
}
