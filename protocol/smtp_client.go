package protos

import (
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"

	"log"
	"strings"
)

const content = `To: gongwenjie09147@hellobike.com
Subject: 您收到一封来自网站的验证邮件
Content-Type: text/html;
X-Mailer: Microsoft OutlookExpress 6.00.2900.2869
<html>
<h1>TEST H1</h1>
</html>
`

func sendMail() {
	// Set up authentication information.
	auth := sasl.NewPlainClient("", "winjeg@qq.com", "xxxxxxxxx")
	to := []string{"amokite@sina.com"}
	msg := strings.NewReader(content)
	err := smtp.SendMail("smtp.qq.com:587", auth, "winjeg@qq.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
