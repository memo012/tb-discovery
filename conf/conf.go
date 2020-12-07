package conf

// Env is discovery env.
type Env struct {
	Region    string
	Zone      string // 机房服务地区标识
	Host      string // 主机标识
	DeployEnv string // 开发环境
}

// Config config.
type Config struct {
	Nodes []string            // 节点信息(ip+port)
	Zones map[string][]string // 多机房服务地区
	Env   *Env                // 环境信息
}
