package EmailSender

import (
	"GoReminder/util"
	"bytes"
	"fmt"
	"log"
	"net/mail"
	"html/template"
	"time"
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
func SendEmail(toEmail string,subject string,content string){
	from := mail.Address{"GoReminder", account}
	to   := mail.Address{"DearUser", toEmail}
	var err error
	if err = client.Mail(from.Address); err != nil {
		util.Appendlog("email_sender,Invaild from email address"+time.Now().String())
	}
	if err = client.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}


	//------------------------------------
	body := content
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k,v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
	message += mime+body

	// Data

	w, err := client.Data()
	if err != nil {
		util.Appendlog("Email56 "+time.Now().String())
	}
	_, err = w.Write([]byte(message))

	if err != nil {
		util.Appendlog("Email61 "+time.Now().String())
	}
	err = w.Close()
	if err != nil {
		util.Appendlog("Email65 "+time.Now().String())
	}
	//client.Quit()
}


func HandleMultipleEmail(){
	channelSize := 100
	emailChan = make(chan string,channelSize)
	ContentChan = make(chan string,channelSize)
	SubJectChan = make(chan string,channelSize)
	var email string
	var content string
	var sub string
	for {
		select {
		 	case email = <-emailChan:
		}
		select {
			case content = <-ContentChan:
		}
		select {
			case sub=<-SubJectChan:
		}
		SendEmail(email, sub,content)
	}
}