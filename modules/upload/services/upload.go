/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-02-21 02:27:31
 * @Description: file content
 */
package services

import (
	"context"
	"fmt"

	"github.com/liziwei01/go-exec-upload/library/logit"
	"github.com/liziwei01/go-exec-upload/library/utils"
	uploadModel "github.com/liziwei01/go-exec-upload/modules/upload/model"
)

// 图片上传
func UploadFile(ctx context.Context, pars uploadModel.UploadPars) error {
	file, err := pars.FileHeader.Open()
	if err != nil {
		err = fmt.Errorf("获取文件字节流失败: %s", err.Error())
		logit.Logger.Error(err)
		return err
	}
	
	return utils.File.SaveFile(file, "data/" + pars.FileHeader.Filename)
}
