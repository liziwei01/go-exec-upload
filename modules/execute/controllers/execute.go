/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:14:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-02-21 02:09:16
 * @Description: file content
 */
package controllers

import (
	"github.com/liziwei01/go-exec-upload/library/response"
	executeModel "github.com/liziwei01/go-exec-upload/modules/execute/model"

	executeService "github.com/liziwei01/go-exec-upload/modules/execute/services"

	"github.com/gin-gonic/gin"
)

func ExecuteLocal(ctx *gin.Context) {
	inputs, hasError := getexecuteLocalPars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	ret, err := executeService.ExecuteLocal(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx, ret)
}

func getexecuteLocalPars(ctx *gin.Context) (executeModel.ExecutePars, bool) {
	var inputs executeModel.ExecutePars

	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}

	return inputs, false
}
