package servers

import (
	"gin-vue-admin/config"

	"github.com/go-gomail/gomail"
)

func SendEmail(to, subject, body string) error {

	m := gomail.NewMessage()

	m.SetHeader("To", to)

	m.SetAddressHeader("From", config.GinVueAdminconfig.SendEmail.FromEmail, "")
	// 主题
	m.SetHeader("Subject", subject)

	// 正文
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer(config.GinVueAdminconfig.SendEmail.ServerHost, config.GinVueAdminconfig.SendEmail.ServerPort, config.GinVueAdminconfig.SendEmail.FromEmail, config.GinVueAdminconfig.SendEmail.FromPasswd)
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
