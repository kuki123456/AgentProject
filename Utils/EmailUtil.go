//// EmailMain
//package Utils
//
//import (
//	"fmt"
//	"net/smtp"
//	"strings"
//)
//
//const(
//	HOST = "smtp.qq.com"
//	SERVER_ADDR = "smtp.qq.com:25"
//	USER = "2278634411@qq.com"
//	PASSWORD = "bqkncqcuxqlleadi"
//)
//
//type Email struct {
//	to string
//	subject string
//	msg string
//}
//
//func NewEmail(to,subject,msg string) *Email{
//	return &Email{to:to,subject:subject,msg:msg}
//}
//
//func SendEmail(email *Email) error{
//	auth := smtp.PlainAuth("",USER,PASSWORD,HOST)
//	sendTo := strings.Split(email.to,";")
//	done := make(chan error,1024)
//
//	go func(){
//		defer close(done)
//		for _,v := range sendTo {
//			//warning ; the last \r\n need twice , not only one .
//			str := strings.Replace("From:"+USER+"~To :"+v+"~Subject:"+email.subject+"~Content-Type: text/plain;charset=UTF-8~","~","\r\n",-1)+"\r\n"+email.msg
//			//fmt.Println("Content:",str)
//			err := smtp.SendMail(SERVER_ADDR,auth,USER,[]string{v},[]byte(str))
//			if err != nil{
//				fmt.Println("Send Error:",err)
//			}
//			done <- err
//		}
//	}()
//
//	for i:=0;i<len(sendTo);i++{
//		<- done
//	}
//
//	return nil
//}
//
//func EmailTo(send string) {
//	email := NewEmail("huqi1@bonree.com","异常统计，请排查！",send)
//	err := SendEmail(email)
//	fmt.Println("result:",err)
//}
package Utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/smtp"
	"strings"
	"time"
)
const (
	emlUser = "2278634411@qq.com"
	emlPwd  = "bqkncqcuxqlleadi"
	emlSMTP = "smtp.qq.com:25"
)
func Eml(ContentText string) error {
	to := "huqi1@bonree.com"
	cc := "huqi1@bonree.com"
	sendTo := strings.Split(to, ";")
	subject := "异常统计邮件,附件为日志，请查看是否异常!"
	boundary := "ds13difsknfsifuere134" //boundary 用于分割邮件内容，可自定义. 注意它的开始和结束格式
	mime := bytes.NewBuffer(nil)
	//设置邮件
	mime.WriteString(fmt.Sprintf("From: %s<%s>\r\nTo: %s\r\nCC: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\n", emlUser, emlUser, to, cc, subject))
	mime.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n", boundary))
	mime.WriteString("Content-Description: 这是一封带附档的邮件\r\n")
	//邮件普通Text正文
	mime.WriteString(fmt.Sprintf("--%s\r\n",boundary))
	mime.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	mime.WriteString(fmt.Sprintf("\r\n%s\r\n",ContentText))
	// 第一个附件
	mime.WriteString(fmt.Sprintf("\n--%s\r\n", boundary))
	mime.WriteString("Content-Type: application/octet-stream\r\n")
	mime.WriteString("Content-Description: 附一个日志文件\r\n")
	mime.WriteString("Content-Transfer-Encoding: base64\r\n")
	mime.WriteString("Content-Disposition: attachment; filename=\"" + fmt.Sprintf("%v%v%v_ERR.log",time.Now().Year(),int(time.Now().Month()),time.Now().Day()) + "\"\r\n\r\n")
	//读取并编码文件内容
	attaData, err := ioutil.ReadFile(fmt.Sprintf("./Log/%v%v%v_ERR.log",time.Now().Year(),int(time.Now().Month()),time.Now().Day()))
	if err != nil {
		return err
	}
	b := make([]byte, base64.StdEncoding.EncodedLen(len(attaData)))
	base64.StdEncoding.Encode(b, attaData)
	mime.Write(b)
	////第二个附件
	//attaData, err := ioutil.ReadFile(fmt.Sprintf("./Log/%v%v%v_ERR.log",time.Now().Year(),int(time.Now().Month()),time.Now().Day()))
	//mime.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	//mime.WriteString("Content-Type: text/plain\r\n")
	//mime.WriteString("Content-Description: 附一个Text文件\r\n")
	//mime.WriteString("Content-Disposition: attachment; filename=\"202158_ERR.log\"\r\n\r\n")
	//mime.WriteString("this is the attachment text") //这里写入的是附件test.txt的内容
	////邮件结束
	//mime.WriteString("\r\n--" + boundary + "--\r\n\r\n")
	//fmt.Println(mime.String())
	//发送相关
	smtpHost, _, err := net.SplitHostPort(emlSMTP)
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", emlUser, emlPwd, smtpHost)
	return smtp.SendMail(emlSMTP, auth, emlUser, sendTo, mime.Bytes())
}