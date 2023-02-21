/*
 * @Author: liziwei01
 * @Date: 2022-04-12 11:15:31
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-02-21 01:55:19
 * @Description: file content
 */
package model

type ExecutePars struct {
	Cmd string `form:"cmd" json:"cmd" binding:"required"`
}
