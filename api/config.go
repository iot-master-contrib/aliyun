package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iot-master-contrib/aliyun/internal"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询阿里云配置
// @Schemes
// @Description 查询阿里云配置
// @Tags config
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[internal.Options] 返回阿里云配置
// @Router /config/aliyun [get]
func configGetAliyun(ctx *gin.Context) {
	curd.OK(ctx, internal.GetOptions())
}

// @Summary 修改阿里云配置
// @Schemes
// @Description 修改阿里云配置
// @Tags config
// @Param cfg body internal.Options true "阿里云配置"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int]
// @Router /config/aliyun [post]
func configSetAliyun(ctx *gin.Context) {
	var conf internal.Options
	err := ctx.BindJSON(&conf)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	internal.SetOptions(conf)
	err = internal.Store()
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	curd.OK(ctx, nil)
}

func configRouter(app *gin.RouterGroup) {
	app.POST("/aliyun", configSetAliyun)
	app.GET("/aliyun", configGetAliyun)
}
