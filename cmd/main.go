package main

import (
	"github.com/iot-master-contrib/aliyun"
	"github.com/iot-master-contrib/aliyun/config"
	"github.com/zgwit/iot-master/v3/pkg/banner"
	"github.com/zgwit/iot-master/v3/pkg/build"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
)

func main() {
	banner.Print("iot-master-plugin:aliyun")
	build.Print()

	config.Load()

	err := log.Open()
	if err != nil {
		log.Fatal(err)
	}

	//加载数据库
	err = db.Open()
	if err != nil {
		log.Fatal(err)
	}

	//MQTT总线
	err = mqtt.Open()
	if err != nil {
		log.Fatal(err)
	}

	app := web.CreateEngine()

	//调用启动
	err = aliyun.Startup(app)
	if err != nil {
		log.Fatal(err)
	}

	err = aliyun.Register()
	if err != nil {
		log.Fatal(err)
	}

	//注册静态页面
	fs := app.FileSystem()
	aliyun.Static(fs)

	//监听HTTP
	app.Serve()
}
