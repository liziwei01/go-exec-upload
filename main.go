/*
 * @Author: liziwei01
 * @Date: 2022-03-03 15:33:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-09-27 22:37:41
 * @Description: main
 */
package main

import (
	"log"

	"github.com/liziwei01/go-exec-upload/bootstrap"
	"github.com/liziwei01/go-exec-upload/httpapi"
)

func main() {

	app, err := bootstrap.Setup()
	if err != nil {
		log.Fatalln(err)
	}
	// 注册接口路由
	httpapi.InitRouters(app.Handler)

	app.Start()
}
