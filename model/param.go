package model

// ArgRegister define registry param.
type ArgRegister struct {
	Env             string   `form:"env" validate:"required"`
	Hostname        string   `form:"hostname" validate:"required"`
	Metadata        string   `form:"metadata"`
	LatestTimestamp int64    `form:"latest_timestamp"`
}
