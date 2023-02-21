/*
 * @Author: liziwei01
 * @Date: 2022-03-20 18:17:39
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-12 22:47:57
 * @Description: file content
 */
package email

import (
	"context"

	"github.com/liziwei01/go-exec-upload/library/logit"

	"gopkg.in/gomail.v2"
)

func (c *client) Send(ctx context.Context, to, subject, body string) error {
	dialer, err := c.connect(ctx)
	if err != nil {
		return err
	}
	msg := gomail.NewMessage()
	msg.SetHeader("From", c.conf.Email.Address)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	err = dialer.DialAndSend(msg)
	if err != nil {
		logit.Logger.Error("send email failed: %s", err.Error())
		return err
	}

	logit.Logger.Info("send email to: %s success", to)

	return nil
}
