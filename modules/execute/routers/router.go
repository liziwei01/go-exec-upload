/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-02-21 02:48:38
 * @Description: file content
 */
package routers

import (
	"github.com/gin-gonic/gin"

	executeController "github.com/liziwei01/go-exec-upload/modules/execute/controllers"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.RouterGroup) {
	executeGroup := router.Group("/api/execute")
	// executeGroup.Use(middleware.CheckLoginMiddleware())
	{
		executeGroup.POST("/exec", executeController.ExecuteLocal)
	}
}
