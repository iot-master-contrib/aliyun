package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(app *gin.RouterGroup) {

	smsRouter(app.Group("/sms"))

	subscriberRouter(app.Group("/subscriber"))

	//configRouter(app.Group("/config"))
}
