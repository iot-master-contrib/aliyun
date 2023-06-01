package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iot-master-contrib/aliyun/types"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询短消息数量
// @Schemes
// @Description 查询短消息数量
// @Tags sms
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回短消息数量
// @Router /sms/count [post]
func noopSmsCount() {}

// @Summary 查询短消息
// @Schemes
// @Description 查询短消息
// @Tags sms
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Sms] 返回短消息信息
// @Router /sms/search [post]
func noopSmsSearch() {}

// @Summary 查询短消息
// @Schemes
// @Description 查询短消息
// @Tags sms
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Sms] 返回短消息信息
// @Router /sms/list [get]
func noopSmsList() {}

// @Summary 创建短消息
// @Schemes
// @Description 创建短消息
// @Tags sms
// @Param search body types.Sms true "短消息信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Sms] 返回短消息信息
// @Router /sms/create [post]
func noopSmsCreate() {}

// @Summary 获取短消息
// @Schemes
// @Description 获取短消息
// @Tags sms
// @Param id path int true "短消息ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Sms] 返回短消息信息
// @Router /sms/{id} [get]
func noopSmsGet() {}

// @Summary 删除短消息
// @Schemes
// @Description 删除短消息
// @Tags sms
// @Param id path int true "短消息ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Sms] 返回短消息信息
// @Router /sms/{id}/delete [get]
func noopSmsDelete() {}

func smsRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Sms]())
	app.POST("/search", curd.ApiSearch[types.Sms]())
	app.GET("/list", curd.ApiList[types.Sms]())
	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.Sms]())
	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDeleteHook[types.Sms](nil, nil))
}
