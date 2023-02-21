/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:15:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-02-21 02:37:18
 * @Description: file content
 */
package model

import "mime/multipart"

type UploadPars struct {
	FileHeader *multipart.FileHeader `form:"file" json:"file"`
}
