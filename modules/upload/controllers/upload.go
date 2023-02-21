/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:14:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-02-21 02:36:55
 * @Description: file content
 */
package controllers

import (
	"fmt"
	"time"

	"github.com/liziwei01/go-exec-upload/library/logit"
	"github.com/liziwei01/go-exec-upload/library/response"
	uploadModel "github.com/liziwei01/go-exec-upload/modules/upload/model"
	uploadService "github.com/liziwei01/go-exec-upload/modules/upload/services"

	"github.com/gin-gonic/gin"
)

func UploadFile(ctx *gin.Context) {
	inputs, hasError := getUploadFilePars(ctx)
	if hasError {
		response.StdInvalidParams(ctx)
		return
	}
	err := uploadService.UploadFile(ctx, inputs)
	if err != nil {
		response.StdFailed(ctx, err.Error())
		return
	}
	response.StdSuccess(ctx)
}

func getUploadFilePars(ctx *gin.Context) (uploadModel.UploadPars, bool) {
	var inputs uploadModel.UploadPars

	startFormat := time.Now()
	err := ctx.ShouldBind(&inputs)
	if err != nil {
		return inputs, true
	}
	costFormat := time.Since(startFormat)

	logit.Logger.Info(fmt.Sprintf("get upload par from request cost=[%s], file bytes is %d", costFormat, inputs.FileHeader.Size))

	return inputs, false
}
