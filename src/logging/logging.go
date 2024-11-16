package logging

/*

References:
https://github.com/fatih/color

*/

import (
	"fmt"
  "time"

	"github.com/fatih/color"
)

// Define colors
var green func(a ...interface{}) string = color.New(color.FgGreen).SprintFunc()
var red func(a ...interface{}) string = color.New(color.FgRed).SprintFunc()
var purple func(a ...interface{}) string = color.New(color.FgMagenta).SprintFunc()
var yellow func(a ...interface{}) string = color.New(color.FgYellow).SprintFunc()
var cyan func(a ...interface{}) string = color.New(color.FgCyan).SprintFunc()
var blue func(a ...interface{}) string = color.New(color.FgBlue).SprintFunc()

// Simple logging functinos with current time and status
func Success(v ...interface{}) {
	fmt.Print(cyan("[") + cyan(GetTime()) + cyan("] ") + green("[SUCCESS] "))
	fmt.Println(v...)
}

func Error(v ...interface{}) {
  fmt.Print(cyan("[") + cyan(GetTime()) + cyan("] ") + red("[ERROR] "))
  fmt.Println(v...)
}

func Info(v ...interface{}) {
  fmt.Print(cyan("[") + cyan(GetTime()) + cyan("] ") + purple("[INFO] "))
  fmt.Println(v...)
}

func Warning(v ...interface{}) {
  fmt.Print(cyan("[") + cyan(GetTime()) + cyan("] ") + yellow("[WARNING] "))
  fmt.Println(v...)
}

func Debug(v ...interface{}) {
  fmt.Print(cyan("[") + cyan(GetTime()) + cyan("] ") + blue("[DEBUG] "))
  fmt.Println(v...)
}

// == Colour functinos ==

func Green(v ...interface{}) {
	green = color.New(color.FgGreen).SprintFunc()
	fmt.Println(green(v...))
}

func Red(v ...interface{}) {
	red = color.New(color.FgRed).SprintFunc()
	fmt.Println(red(v...))
}

func Purple(v ...interface{}) {
	purple = color.New(color.FgMagenta).SprintFunc()
	fmt.Println(purple(v...))
}

func Blue(v ...interface{}) {
	blue = color.New(color.FgBlue).SprintFunc()
	fmt.Println(blue(v...))
}

func Cyan(v ...interface{}) {
	cyan = color.New(color.FgCyan).SprintFunc()
	fmt.Println(cyan(v...))
}

// == The following functions return the coloured input ==

func SGreen(v ...interface{}) string {
	green = color.New(color.FgGreen).SprintFunc()
	return fmt.Sprint(green(v...))
}

func SRed(v ...interface{}) string {
	red = color.New(color.FgRed).SprintFunc()
	return fmt.Sprint(red(v...))
}

func SPurple(v ...interface{}) string {
	purple = color.New(color.FgMagenta).SprintFunc()
	return fmt.Sprint(purple(v...))
}

func SBlue(v ...interface{}) string {
	blue = color.New(color.FgBlue).SprintFunc()
	return fmt.Sprint(blue(v...))
}

func SCyan(v ...interface{}) string {
	cyan = color.New(color.FgCyan).SprintFunc()
	return fmt.Sprint(cyan(v...))
}

// Function to get current time
func GetTime() string {
	return time.Now().Format("15:04:05")
}

