/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-02-21 02:12:49
 * @Description: file content
 */
package routers

import (
	"github.com/gin-gonic/gin"

	uploadController "github.com/liziwei01/go-exec-upload/modules/upload/controllers"
)

/**
 * @description: 后台路由分发
 * @param {*}
 * @return {*}
 */
func Init(router *gin.RouterGroup) {
	uploadGroup := router.Group("/api/upload")
	{
		uploadGroup.POST("/file", uploadController.UploadFile)
	}
}
