package logging

import (
  "fmt"
  "github.com/fatih/color"
)

var green func(a ...interface{}) string
var red func(a ...interface{}) string
var purple func(a ...interface{}) string
var yellow func(a ...interface{}) string // Define colors
var cyan func(a ...interface{}) string
var blue func(a ...interface{}) string

// Println functions
func Goodln(str_to_print string){
  green = color.New(color.FgGreen).SprintFunc()
  fmt.Println(green("[+] ") + str_to_print)
}

func Badln(str_to_print string){
  red = color.New(color.FgRed).SprintFunc()
  fmt.Println(red("[-] ") + str_to_print)
}

func Infoln(str_to_print string){
  purple = color.New(color.FgMagenta).SprintFunc()
  fmt.Println(purple("[*] ") + str_to_print)
}

func Greenln(str_to_print string){
  green = color.New(color.FgGreen).SprintFunc()
  fmt.Println(green(str_to_print))
}

func Redln(str_to_print string){
  red = color.New(color.FgRed).SprintFunc()
  fmt.Println(red(str_to_print))
}

func Purpleln(str_to_print string){
  purple = color.New(color.FgMagenta).SprintFunc()
  fmt.Println(purple(str_to_print))
}

func Blueln(str_to_print string){
  blue = color.New(color.FgBlue).SprintFunc()
  fmt.Println(blue(str_to_print))
}

func Cyanln(str_to_print string){
  cyan = color.New(color.FgCyan).SprintFunc()
  fmt.Println(cyan(str_to_print))
}

// == Separator between Println and Print functions ==

func Good(str_to_print string){
  green = color.New(color.FgCyan).SprintFunc()
  fmt.Print(green("[+] ") + str_to_print)
}

func Bad(str_to_print string){
  red = color.New(color.FgRed).SprintFunc()
  fmt.Print(red("[-] ") + str_to_print)
}

func Info(str_to_print string){
  purple = color.New(color.FgMagenta).SprintFunc()
  fmt.Print(purple("[*] ") + str_to_print)
}

func Green(str_to_print string){
  green = color.New(color.FgGreen).SprintFunc()
  fmt.Print(green(str_to_print))
}

func Red(str_to_print string){
  red = color.New(color.FgRed).SprintFunc()
  fmt.Print(red(str_to_print))
}

func Purple(str_to_print string){
  purple = color.New(color.FgMagenta).SprintFunc()
  fmt.Print(purple(str_to_print))
}

func Blue(str_to_print string){
  blue = color.New(color.FgBlue).SprintFunc()
  fmt.Print(blue(str_to_print))
}

func Cyan(str_to_print string){
  cyan = color.New(color.FgCyan).SprintFunc()
  fmt.Print(cyan(str_to_print))
}

// == This functions return the coloured input ==

func SGreen(str_to_print string) (string) {
  green = color.New(color.FgGreen).SprintFunc()
  return fmt.Sprint(green(str_to_print))
}

func SRed(str_to_print string) (string) {
  red = color.New(color.FgRed).SprintFunc()
  return fmt.Sprint(red(str_to_print))
}

func SPurple(str_to_print string) (string) {
  purple = color.New(color.FgMagenta).SprintFunc()
  return fmt.Sprint(purple(str_to_print))
}

func SBlue(str_to_print string) (string) {
  blue = color.New(color.FgBlue).SprintFunc()
  return fmt.Sprint(blue(str_to_print))
}

func SCyan(str_to_print string) (string) {
  cyan = color.New(color.FgCyan).SprintFunc()
  return fmt.Sprint(cyan(str_to_print))
}




