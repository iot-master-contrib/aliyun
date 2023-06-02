package aliyun

import (
	"embed"
	"encoding/json"
	"github.com/iot-master-contrib/aliyun/api"
	_ "github.com/iot-master-contrib/aliyun/docs"
	"github.com/iot-master-contrib/aliyun/types"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"net/http"
)

func App() *model.App {
	return &model.App{
		Id:   "sms",
		Name: "短消息推送",
		Entries: []model.AppEntry{{
			Path: "app/aliyun/subscriber",
			Name: "消息订阅",
		}, {
			Path: "app/aliyun/sms",
			Name: "消息历史",
		}, {
			Path: "app/aliyun/setting",
			Name: "配置",
		}},
		Type:    "tcp",
		Address: "http://localhost" + web.GetOptions().Addr,
		Icon:    "/app/sms/assets/sms.svg",
	}
}

//go:embed all:app/sms
var wwwFiles embed.FS

// @title 阿里云接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /app/aliyun/api/
// @query.collection.format multi
func main() {
}

func Startup(app *web.Engine) error {

	//同步表结构
	err := db.Engine.Sync2(
		new(types.Sms),
	)
	if err != nil {
		log.Fatal(err)
	}

	//注册前端接口
	api.RegisterRoutes(app.Group("/app/sms/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(app.Group("/app/sms"), "sms")

	return nil
}

func Register() error {
	payload, _ := json.Marshal(App())
	return mqtt.Publish("master/register", payload, false, 0)
}

func Static(fs *web.FileSystem) {
	//前端静态文件
	fs.Put("/app/sms", http.FS(wwwFiles), "", "app/sms/index.html")
}

func Shutdown() error {

	//只关闭Web就行了，其他通过defer关闭

	return nil
}
