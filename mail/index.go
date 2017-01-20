package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

//-------------------------------------
//
//
//
//-------------------------------------
func Send(host, port, user, pass string, mime string, title, body string, tos []string) error {
	dest := fmt.Sprintf("%s:%s", host, port)
	conn, err := tls.Dial("tcp", dest, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return err
	}

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	defer client.Close()

	if err := client.Auth(smtp.PlainAuth(
		"",
		user,
		pass,
		host,
	)); err != nil {
		return err
	}

	if err := client.Mail(user); err != nil {
		return err
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
			return err
		}
	}

	//
	w, err := client.Data()
	if err != nil {
		return err
	}
	w.Write([]byte(message))
	w.Close()
	return nil
}

//-------------------------------------
//
//
//
//-------------------------------------
func SendHtml(host, port, user, pass, title, body string, tos []string) error {
	return Send(host, port, user, pass, `text/html; chartset="utf-8"`, title, body, tos)
}
