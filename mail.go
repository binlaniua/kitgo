package kitgo

import (
	"net/smtp"
	"log"
	"strings"
	"fmt"
	"crypto/tls"
)

//-------------------------------------
//
// 
//
//-------------------------------------
func MailSend(host, port, user, pass string, mime string, title, body string, tos []string) bool {
	dest := fmt.Sprintf("%s:%s", host, port)
	conn, err := tls.Dial("tcp", dest, &tls.Config{
		InsecureSkipVerify:true,
	})
	if err != nil {
		log.Println("发送邮件, 创建连接失败 => ", err)
		return false
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Println("发送邮件, 创建客户端失败 => ", err)
		return false
	}
	defer client.Close()

	if err := client.Auth(smtp.PlainAuth(
		"",
		user,
		pass,
		host,
	)); err != nil {
		log.Println("发送邮件, 用户名密码错误 => ", err)
		return false
	}

	if err := client.Mail(user); err != nil {
		log.Println("发送邮件, 设置发件人错误 => ", err)
		return false
	}

	//
	header := make(map[string]string)
	header["From"] = user
	header["To"] = strings.Join(tos, ";")
	header["Subject"] = title
	header["Content-Type"] = mime
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	//
	for _, to := range tos {
		if err = client.Rcpt(to); err != nil {
			log.Println("发送邮件, 设置收件人错误 => ", err)
			return false
		}
	}

	//
	w, err := client.Data()
	w.Write([]byte(message))
	w.Close()
	return true
}

//-------------------------------------
//
// 
//
//-------------------------------------
func MailSendHtml(host, port, user, pass, title, body string, tos []string) bool {
	return MailSend(host, port, user, pass, `text/html; chartset="utf-8"`, title, body, tos)
}
