package util

import "os"
var file *os.File
func init(){
	const LOGFILE="log.txt"
	os.Remove(LOGFILE)
	file,_ =os.OpenFile(LOGFILE,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
}
func Appendlog(logtext string)  {
	file.Write([]byte(logtext))
}