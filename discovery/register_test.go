package discovery

import (
	"fmt"
	dc "github.com/memo012/tb-discovery/conf"
	"github.com/memo012/tb-discovery/model"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var (
	fet    = &model.ArgFetch{AppID: "main.arch.test", Zone: "sh001", Env: "pre"}
)

var (
	reg = defRegisArg()
)

var config = newConfig()

func defRegisArg() *model.ArgRegister {
	return &model.ArgRegister{
		AppID:           "main.arch.test",                // 服务唯一标识
		Hostname:        "test1",                         // 主机名
		Zone:            "sh001",                         // 机房服务标识
		Env:             "pre",                           // 环境
		Metadata:        `{"test":"test","weight":"10"}`, // 扩展元数据
		LatestTimestamp: time.Now().UnixNano(),           // 时间戳
	}
}

func newConfig() *dc.Config {
	c := &dc.Config{
		Nodes: []string{"127.0.0.1:7171", "127.0.0.1:7172"},
		Env: &dc.Env{
			Zone:      "tb001",
			DeployEnv: "pre",
			Host:      "test_server",
		},
	}
	return c
}

func TestDiscovery_Register(t *testing.T) {
	Convey("test Register", t, func() {
		svr, cancel := New(config)
		defer cancel()
		i := model.NewInstance(reg)
		svr.Register(i, reg.LatestTimestamp)
		// 拉取 instance
		ins, err := svr.Fetch(fet)
		fmt.Println(ins)
		So(err, ShouldBeNil)
		So(len(ins.Instances), ShouldResemble, 1)
	})
}
