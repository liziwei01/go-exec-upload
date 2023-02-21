/*
 * @Author: liziwei01
 * @Date: 2022-03-04 15:40:52
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 14:48:10
 * @Description: file content
 */
package logit

import (
	"context"

	lib "github.com/baidu/go-lib/log"
	"github.com/baidu/go-lib/log/log4go"
)

var (
	Logger      log4go.Logger
	levelStr    = "INFO"
	logDir      = "./log"
	hasStdOut   = true
	when        = "H"
	backupCount = 5
)

/**
 * @description: all the log are recorded under ./log
 * @param {string} programName
 * @return {*}
 */
func Init(ctx context.Context, programName string) error {
	var err error
	Logger, err = lib.Create(programName, levelStr, logDir, hasStdOut, when, backupCount)
	return err
}