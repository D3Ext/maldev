package all

import "github.com/D3Ext/maldev/logging"

// Println functions
func Goodln(v ...interface{}){
  logging.Goodln(v...)
}

func Badln(v ...interface{}){
  logging.Badln(v...)
}

func Infoln(v ...interface{}){
  logging.Infoln(v...)
}

func Greenln(v ...interface{}){
  logging.Greenln(v...)
}

func Redln(v ...interface{}){
  logging.Redln(v...)
}

func Purpleln(v ...interface{}){
  logging.Purpleln(v...)
}

func Blueln(v ...interface{}){
  logging.Blueln(v...)
}

func Cyanln(v ...interface{}){
  logging.Cyanln(v...)
}

// == Separator between Println and Print functions ==

func Good(v ...interface{}){
  logging.Good(v...)
}

func Bad(v ...interface{}){
  logging.Bad(v...)
}

func Info(v ...interface{}){
  logging.Info(v...)
}

func Green(v ...interface{}){
  logging.Green(v...)
}

func Red(v ...interface{}){
  logging.Red(v...)
}

func Purple(v ...interface{}){
  logging.Purple(v...)
}

func Blue(v ...interface{}){
  logging.Blue(v...)
}

func Cyan(v ...interface{}){
  logging.Cyan(v...)
}

// == This functions return the coloured input ==

func SGreen(v ...interface{}) (string) {
  return logging.SGreen(v)
}

func SRed(v ...interface{}) (string) {
  return logging.SRed(v)
}

func SPurple(v ...interface{}) (string) {
  return logging.SPurple(v)
}

func SBlue(v ...interface{}) (string) {
  return logging.SBlue(v)
}

func SCyan(v ...interface{}) (string) {
  return logging.SCyan(v)
}

// == This functions return or print ASCII banners ==

func GetBanner(str_to_convert string) (string) {
  return logging.GetBanner(str_to_convert)
}

func PrintBanner(str_to_convert string) {
  logging.PrintBanner(str_to_convert)
}

// == Time related functions ==

func GetTime() (string) {
  return logging.GetTime()
}

func TimePrintln(input string) {
  logging.TimePrintln(input)
}

func TimePrint(input string) {
  logging.TimePrint(input)
}

func TimeSprint(input string) (string) {
  return logging.TimeSprint(input)
}

