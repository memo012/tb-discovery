package model

import "sync"

type Instance struct {
	Env string `json:"env"`
	HostName string `json:"hostname"`
	Metadata map[string]string `json:"metadata"`

	// timestamp
	RegTimestamp   int64 `json:"reg_timestamp"`
	LatestTimestamp int64 `json:"latest_timestamp"` // 服务最新更新时间
}

type App struct {
	lock sync.RWMutex
}


// NewInstance new a instance.
func (a *App) NewInstance(ins *Instance, latestTime int64) (i *Instance, ok bool)  {

	return nil, false
}
