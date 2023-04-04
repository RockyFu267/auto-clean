package controller

//EnvConf yaml配置的key不能用 _ 否则会信息丢失
type EnvConf struct {
	Interval int64         `json:"interval"`
	Service  []ServiceInfo `json:"service"`
}

type ServiceInfo struct {
	Cyclelength string `json:"cyclelength"`
	Cleanpath   string `json:"cleanpath"`
	Cleanname   string `json:"cleanname"`
}
