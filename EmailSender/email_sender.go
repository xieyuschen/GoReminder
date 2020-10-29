package EmailSender

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
func SendEmail(toEmail string,subject string,content string){
	from := mail.Address{"GoReminder", account}
	to   := mail.Address{"DearUser", toEmail}
	var err error
	if err = client.Mail(from.Address); err != nil {
		log.Panic(err)
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