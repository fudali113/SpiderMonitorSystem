package models

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
)

const (
	user     = "fuyi@shanlaohu.com"
	password = ""
	host     = "smtp.exmail.qq.com:25"
)

var (
	ToAddress = "591327191@qq.com"
)

type Email struct {
	To       string
	Subject  string
	Body     string
	MailType string
}

func SendEmail(email Email) error {
	to := email.To
	subject := email.Subject
	body := email.Body
	mailtype := email.MailType
	return SendToMail(user, password, host, to, subject, body, mailtype)
}
func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func GetHtmlWithTpl(file string, m map[string]interface{}) (string, error) {
	var body bytes.Buffer
	t, err := template.ParseFiles(file)
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(&body, m)
	}

	bodyStr := string(body.Bytes())
	return bodyStr, err
}

func main() {
	to := "591327191@qq.com"

	subject := "使用Golang发送邮件"
	var body bytes.Buffer
	t, err := template.ParseFiles("../views/email.tpl")
	if err != nil {
		fmt.Println(err)
	} else {
		m := map[string]string{
			"before":   "1",
			"pc_id":    "1",
			"lastData": "1"}
		t.Execute(&body, m)
	}

	bodyStr := string(body.Bytes())

	email := Email{To: to, Subject: subject, Body: bodyStr, MailType: "html"}
	fmt.Println("send email")
	err1 := SendEmail(email)
	if err1 != nil {
		fmt.Println("error")
	}

}
