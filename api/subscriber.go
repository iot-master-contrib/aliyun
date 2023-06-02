package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iot-master-contrib/aliyun/types"
	"github.com/zgwit/iot-master/v3/pkg/curd"
)

// @Summary 查询订阅数量
// @Schemes
// @Description 查询订阅数量
// @Tags subscriber
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回订阅数量
// @Router /subscriber/count [post]
func noopSubscriberCount() {}

// @Summary 查询订阅
// @Schemes
// @Description 查询订阅
// @Tags subscriber
// @Param search body ParamSearch true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Subscriber] 返回订阅信息
// @Router /subscriber/search [post]
func noopSubscriberSearch() {}

// @Summary 查询订阅
// @Schemes
// @Description 查询订阅
// @Tags subscriber
// @Param search query ParamList true "查询参数"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyList[types.Subscriber] 返回订阅信息
// @Router /subscriber/list [get]
func noopSubscriberList() {}

// @Summary 创建订阅
// @Schemes
// @Description 创建订阅
// @Tags subscriber
// @Param search body types.Subscriber true "订阅信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Subscriber] 返回订阅信息
// @Router /subscriber/create [post]
func noopSubscriberCreate() {}

// @Summary 修改订阅
// @Schemes
// @Description 修改订阅
// @Tags subscriber
// @Param id path int true "订阅ID"
// @Param subscriber body types.Subscriber true "订阅信息"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Subscriber] 返回订阅信息
// @Router /subscriber/{id} [post]
func noopSubscriberUpdate() {}

// @Summary 获取订阅
// @Schemes
// @Description 获取订阅
// @Tags subscriber
// @Param id path int true "订阅ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Subscriber] 返回订阅信息
// @Router /subscriber/{id} [get]
func noopSubscriberGet() {}

// @Summary 删除订阅
// @Schemes
// @Description 删除订阅
// @Tags subscriber
// @Param id path int true "订阅ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Subscriber] 返回订阅信息
// @Router /subscriber/{id}/delete [get]
func noopSubscriberDelete() {}

// @Summary 启用订阅
// @Schemes
// @Description 启用订阅
// @Tags subscriber
// @Param id path int true "订阅ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Subscriber] 返回订阅信息
// @Router /subscriber/{id}/enable [get]
func noopSubscriberEnable() {}

// @Summary 禁用订阅
// @Schemes
// @Description 禁用订阅
// @Tags subscriber
// @Param id path int true "订阅ID"
// @Accept json
// @Produce json
// @Success 200 {object} ReplyData[types.Subscriber] 返回订阅信息
// @Router /subscriber/{id}/disable [get]
func noopSubscriberDisable() {}

// @Summary 导出订阅
// @Schemes
// @Description 导出订阅
// @Tags product
// @Accept json
// @Produce octet-stream
// @Router /subscriber/export [get]
func noopSubscriberExport() {}

// @Summary 导入订阅
// @Schemes
// @Description 导入订阅
// @Tags product
// @Param file formData file true "压缩包"
// @Accept mpfd
// @Produce json
// @Success 200 {object} ReplyData[int64] 返回订阅数量
// @Router /subscriber/import [post]
func noopSubscriberImport() {}

func subscriberRouter(app *gin.RouterGroup) {

	app.POST("/count", curd.ApiCount[types.Subscriber]())
	app.POST("/search", curd.ApiSearch[types.Subscriber]())
	app.GET("/list", curd.ApiList[types.Subscriber]())
	app.POST("/create", curd.ApiCreateHook[types.Subscriber](curd.GenerateRandomId[types.Subscriber](8), nil))
	app.GET("/:id", curd.ParseParamId, curd.ApiGet[types.Subscriber]())
	app.POST("/:id", curd.ParseParamId, curd.ApiUpdateHook[types.Subscriber](nil, nil,
		"id", "user", "cellphone", "level", "disabled"))
	app.GET("/:id/delete", curd.ParseParamId, curd.ApiDeleteHook[types.Subscriber](nil, nil))
	app.GET("/export", curd.ApiExport("subscriber", "订阅人"))
	app.POST("/import", curd.ApiImport("subscriber"))

	app.GET(":id/disable", curd.ParseParamId, curd.ApiDisableHook[types.Subscriber](true, nil, nil))
	app.GET(":id/enable", curd.ParseParamId, curd.ApiDisableHook[types.Subscriber](false, nil, nil))
}
