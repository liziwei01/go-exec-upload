/*
 * @Author: liziwei01
 * @Date: 2022-04-12 10:45:14
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-02-21 02:09:09
 * @Description: file content
 */
package services

import (
	"context"
	"os"
	"os/exec"
	"strings"

	executeModel "github.com/liziwei01/go-exec-upload/modules/execute/model"
)

func ExecuteLocal(ctx context.Context, pars executeModel.ExecutePars) (string, error) {
	cmdString := strings.Split(pars.Cmd, " ")
	cmd := exec.Command(cmdString[0], cmdString[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return "", err
}
