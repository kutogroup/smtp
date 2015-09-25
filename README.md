# smtp
Golang smtp

Example:

```go

package main

import (
	"github.com/MouseSun/goprint"
	"github.com/MouseSun/smtp"
)

func main() {

	m_smtp := smtp.SmtpHelper{
		"xxx@xxx.com",					//user
		"xxxxxxx",						//password
		"smtp.xxxx.com:25",				//host:port
		"sunzhigang@maiyamall.com",		//receiver email address(more than one, divider by ';')
		"subject",						//subject
		"body",							//body
		false,							//ishtml
	}

	goprint.P("smtp", m_smtp)
	if err := m_smtp.SendMail(); err != nil {
		goprint.V("Send Mail failed")
	} else {
		goprint.V("Send Mail success")
	}
}

```
