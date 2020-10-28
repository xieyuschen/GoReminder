package services

import (
	"GoReminder/models"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"net"
	"net/smtp"
)
//#######################################################
//--------------database services---------------------------

//-------------------------------------------------------

//#######################################################
//--------------email services---------------------------
//The following variables are defined for email services
var client *smtp.Client
var account string
var password string
var emailChan chan string
var ContentChan chan string
var SubJectChan chan string
//-------------------------------------------------------

//Case Sensitive! Need a uppercase suffix


func init(){

	emailInit()
}

func ReadSettingsFromFile(settingFilePath string) (config models.Config){
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}
func emailInit(){

	conf := ReadSettingsFromFile("Config.json")
	account =conf.EmailSenderSettings.Email
	password = conf.EmailSenderSettings.Password
	// Connect to the SMTP Server
	servername := "smtp.163.com:465"

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("",account, password, host)

	// TLS config
	tlsconfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		//log.Panic(err)
		fmt.Println(err)
	}

	client, err = smtp.NewClient(conn, host)
	if err != nil {
		//log.Panic(err)
		fmt.Println(err)

	}

	// Auth
	if err = client.Auth(auth); err != nil {
		//log.Panic(err)
		fmt.Println(err)
	}
	go HandleMultipleEmail()

}

