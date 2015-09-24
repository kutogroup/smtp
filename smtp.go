package smtp

import (
	"net/smtp"
	"strings"
)

type SmtpHelper struct {
	UserName string
	PassWord string
	HostName string
	SendTo   string
	Subject  string
	Body     string
	IsHtml   bool
}

func (helper *SmtpHelper) SendMail() error {
	hp := strings.Split(helper.HostName, ":")
	auth := smtp.PlainAuth("", helper.UserName, helper.PassWord, hp[0])
	var content_type string
	if helper.IsHtml == true {
		content_type = "Content-Type: text/html; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain; charset=UTF-8"
	}

	msg := []byte("To: " + helper.SendTo + "\r\nFrom: " + helper.UserName + "<" + helper.UserName +
		">\r\nSubject: " + helper.Subject + "\r\n" + content_type + "\r\n\r\n" + helper.Body)
	send_to := strings.Split(helper.SendTo, ";")
	err := smtp.SendMail(helper.HostName, auth, helper.UserName, send_to, msg)
	return err
}
