package services

import (

	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"os"

	"net"
	"net/smtp"
)
//#######################################################
//--------------database services---------------------------
var db *gorm.DB

//-------------------------------------------------------

//#######################################################
//--------------email services---------------------------
//The following variables are defined for email services
var client *smtp.Client
var account string
var password string
var emailChan chan string
var verifyCodeChan chan string
//-------------------------------------------------------

type Config struct {
	DbSettings DbSettings `json:"DbSettings"`
	EmailSenderSettings EmailSenderSettings `json:"EmailSenderSettings"`
}
type DbSettings struct{
	Username string `json:"Username"`
	Password string	`json:"Password"`
	Hostname string `json:"Hostname"`
	Dbname   string `json:"Dbname"`
}

//Case Sensitive! Need a uppercase suffix
type EmailSenderSettings struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func init(){
	databaseInit()
	emailInit()
}
func databaseInit(){
	settings := DbSettings{Username: "root", Password: "root", Hostname: "127.0.0.1:3306", Dbname: "husthole"}
	//"root:@tcp(127.0.0.1:3306)/?parseTime=true&charset=utf8"
	connStr := dsn(settings)
	msdb, _ := sql.Open("mysql",connStr)

	msdb.Exec("create database if not exists "+settings.Dbname +" character set utf8")
	msdb.Close()

	db, _ = gorm.Open("mysql",dsn(settings))
	//db.DB().SetMaxIdleConns(0)

}
func ReadSettingsFromFile(settingFilePath string) (config Config){
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
func dsn(settings DbSettings) string {
	// https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	// Add ?parseTime=true
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8", settings.Username,settings.Password, settings.Hostname,settings.Dbname)
}
