package models

import (
	"bytes"
	"html/template"
	"net/smtp"
	"strings"

	"github.com/astaxie/beego"
)

const (
	DefaultTA = "591327191@qq.com"
)

var (
	user      = beego.AppConfig.String("email.user")
	password  = beego.AppConfig.String("email.passwd")
	host      = beego.AppConfig.String("email.host")
	ToAddress = DefaultTA
)

type Email struct {
	To       string
	Subject  string
	Body     string
	MailType string
}

func SendEmail(email Email) {
	to := email.To
	subject := email.Subject
	body := email.Body
	mailtype := email.MailType
	err := SendToMail(user, password, host, to, subject, body, mailtype)
	if err != nil {
		beego.Error(err)
		SendEmail(email)
	}
}
func SendEmailWithMap(m map[string]interface{}, sub string, tpl string) {
	body, _ := GetHtmlWithTpl(tpl, m)
	email := Email{To: ToAddress,
		Subject:  sub,
		Body:     body,
		MailType: "html"}

	SendEmail(email)
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
	if err == nil {
		beego.Notice("send one email to : ", to)
	}
	return err
}

func GetHtmlWithTpl(file string, m map[string]interface{}) (string, error) {
	var body bytes.Buffer
	t, err := template.ParseFiles(file)
	if err != nil {
		beego.Error(err)
	} else {
		t.Execute(&body, m)
	}

	bodyStr := string(body.Bytes())
	return bodyStr, err
}

//func main() {
//	to := "591327191@qq.com"

//	subject := "使用Golang发送邮件"
//	var body bytes.Buffer
//	t, err := template.ParseFiles("../views/email.tpl")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		m := map[string]string{
//			"before":   "1",
//			"pc_id":    "1",
//			"lastData": "1"}
//		t.Execute(&body, m)
//	}

//	bodyStr := string(body.Bytes())

//	email := Email{To: to, Subject: subject, Body: bodyStr, MailType: "html"}
//	beego.Error("send email")
//	SendEmail(email)

//}
