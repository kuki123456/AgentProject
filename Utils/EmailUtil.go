// EmailMain
package Utils

import (
	"fmt"
	"net/smtp"
	"strings"
)

const(
	HOST = "smtp.qq.com"
	SERVER_ADDR = "smtp.qq.com:25"
	USER = "2278634411@qq.com"
	PASSWORD = "bqkncqcuxqlleadi"
)

type Email struct {
	to string
	subject string
	msg string
}

func NewEmail(to,subject,msg string) *Email{
	return &Email{to:to,subject:subject,msg:msg}
}

func SendEmail(email *Email) error{
	auth := smtp.PlainAuth("",USER,PASSWORD,HOST)
	sendTo := strings.Split(email.to,";")
	done := make(chan error,1024)

	go func(){
		defer close(done)
		for _,v := range sendTo {
			//warning ; the last \r\n need twice , not only one .
			str := strings.Replace("From:"+USER+"~To :"+v+"~Subject:"+email.subject+"~Content-Type: text/plain;charset=UTF-8~","~","\r\n",-1)+"\r\n"+email.msg
			//fmt.Println("Content:",str)
			err := smtp.SendMail(SERVER_ADDR,auth,USER,[]string{v},[]byte(str))
			if err != nil{
				fmt.Println("Send Error:",err)
			}
			done <- err
		}
	}()

	for i:=0;i<len(sendTo);i++{
		<- done
	}

	return nil
}

func EmailTo(send string) {
	email := NewEmail("huqi1@bonree.com","异常统计，请排查！",send)
	err := SendEmail(email)
	fmt.Println("result:",err)
}
