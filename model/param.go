package model

// ArgRegister define registry param.
type ArgRegister struct {
	Zone            string `json:"zone"`
	AppID           string `json:"appid"` // 服务名标识
	Env             string `form:"env" validate:"required"`
	Hostname        string `form:"hostname" validate:"required"`
	Metadata        string `form:"metadata"`
	LatestTimestamp int64  `form:"latest_timestamp"`
}

// ArgFetch define fetch param.
type ArgFetch struct {
	Zone   string `form:"zone"`
	Env    string `form:"env" validate:"required"`
	AppID  string `form:"appid" validate:"required"`
}