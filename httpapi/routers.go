/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:04:46
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-02-21 01:52:01
 * @Description: 路由分发
 */

package httpapi

import (
	"net/http"

	"github.com/liziwei01/go-exec-upload/middleware"
	executeRouters "github.com/liziwei01/go-exec-upload/modules/execute/routers"

	uploadRouters "github.com/liziwei01/go-exec-upload/modules/upload/routers"

	"github.com/gin-gonic/gin"
)

/**
 * @description: start http server and start listening
 * @param {*}
 * @return {*}
 */
func InitRouters(handler *gin.Engine) {
	//暂时解决跨域问题
	handler.Use(middleware.CrossRegionMiddleware())
	// router.Use(middleware.CheckTokenMiddleware(), middleware.GetFrequencyControlMiddleware(), middleware.PostFrequencyControlMiddleware(), middleware.MailFrequencyControlMiddleware())
	// init routers
	router := handler.Group("/")

	uploadRouters.Init(router)
	executeRouters.Init(router)

	// safe router
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello! THis is iDownloader. Welcome to our offical website!")
	})
}
