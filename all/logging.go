package all

import "github.com/D3Ext/maldev/src/logging"

func Success(v ...interface{}) {
  logging.Success(v...)
}

func Error(v ...interface{}) {
  logging.Error(v...)
}

func Info(v ...interface{}) {
  logging.Info(v...)
}

func Warning(v ...interface{}) {
  logging.Warning(v...)
}

func Debug(v ...interface{}) {
  logging.Debug(v...)
}

// Colours functions

func Green(v ...interface{}) {
	logging.Green(v...)
}

func Red(v ...interface{}) {
	logging.Red(v...)
}

func Purple(v ...interface{}) {
	logging.Purple(v...)
}

func Blue(v ...interface{}) {
	logging.Blue(v...)
}

func Cyan(v ...interface{}) {
	logging.Cyan(v...)
}

// == This functions return the coloured input ==

func SGreen(v ...interface{}) string {
	return logging.SGreen(v)
}

func SRed(v ...interface{}) string {
	return logging.SRed(v)
}

func SPurple(v ...interface{}) string {
	return logging.SPurple(v)
}

func SBlue(v ...interface{}) string {
	return logging.SBlue(v)
}

func SCyan(v ...interface{}) string {
	return logging.SCyan(v)
}

// == This functions return or print ASCII banners ==

func GetBanner(str_to_convert string) string {
	return logging.GetBanner(str_to_convert)
}

func PrintBanner(str_to_convert string) {
	logging.PrintBanner(str_to_convert)
}

func GetTime() string {
	return logging.GetTime()
}

