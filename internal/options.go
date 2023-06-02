package internal

import (
	"github.com/zgwit/iot-master/v3/pkg/config"
	"github.com/zgwit/iot-master/v3/pkg/env"
)

type Options struct {
	Id       string `json:"id"`
	Secret   string `json:"secret"`
	Sign     string `json:"sign"`
	Template string `json:"template"`
}

func Default() Options {
	return Options{}
}

var options Options = Default()
var configure = config.AppName() + ".aliyun.yaml"

const ENV = "IOT_MASTER_ALIYUN_"

func GetOptions() Options {
	return options
}

func SetOptions(opts Options) {
	options = opts
}

func init() {
	options.FromEnv()
}

func (options *Options) FromEnv() {
	options.Id = env.Get(ENV+"ID", options.Id)
	options.Secret = env.Get(ENV+"SECRET", options.Secret)
	options.Sign = env.Get(ENV+"SIGN", options.Sign)
	options.Template = env.Get(ENV+"TEMPLATE", options.Template)
}

func (options *Options) ToEnv() []string {
	ret := []string{
		ENV + "ID=" + options.Id,
		ENV + "SECRET=" + options.Secret,
		ENV + "SIGN=" + options.Sign,
		ENV + "TEMPLATE=" + options.Template,
	}
	return ret
}

func Load() error {
	return config.Load(configure, &options)
}

func Store() error {
	return config.Store(configure, &options)
}
