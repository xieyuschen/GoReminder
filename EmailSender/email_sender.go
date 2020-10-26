package services

import (
	"bytes"
	"fmt"
	"log"
	"net/mail"
	"html/template"
)
func parseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}

	return nil
}
func SendEmail(toEmail string,verifycode string){
	from := mail.Address{"HoleTeam", account}
	to   := mail.Address{"DearUser", toEmail}
	var err error
	if err = client.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = client.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	subj := "HustHole注册"

	content := "欢迎来到华科树洞，你的验证码为："+verifycode+" 验证码有效期20分钟，如果不是您本人注册，请忽略此信息"




	//===================================
	//Send a email template
	t, err := template.ParseFiles("1.html")
	if err != nil {
		log.Panic(err)
	}
	buffer := new(bytes.Buffer)
	var data interface{}
	if err = t.Execute(buffer, data); err != nil {
		log.Panic(err)
	}
	content = buffer.String()

	//------------------------------------
	body := content
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k,v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
	message += mime + body

	// Data

	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write([]byte(message))

	if err != nil {
		log.Panic(err)
	}
	err = w.Close()
	if err != nil {
		log.Panic(err)
	}
	//client.Quit()
}

func AddEmailVerifyCodeToChennel(toEmail string,verifycode string){
	emailChan<-toEmail
	verifyCodeChan<-verifycode
}

func HandleMultipleEmail(){
	channelSize := 100
	emailChan = make(chan string,channelSize)
	verifyCodeChan = make(chan string,channelSize)
	var email string
	var code string
	for {
		select {
		 	case email = <-emailChan:
		}
		select {
			case code = <-verifyCodeChan:
		}
		SendEmail(email,code)
	}
}