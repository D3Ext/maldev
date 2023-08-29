package logging

/*

References:
https://github.com/fatih/color

*/

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
)

var green func(a ...interface{}) string
var red func(a ...interface{}) string
var purple func(a ...interface{}) string
var yellow func(a ...interface{}) string // Define colors
var cyan func(a ...interface{}) string
var blue func(a ...interface{}) string

// Println functions
func Goodln(v ...interface{}) {
	if runtime.GOOS == "linux" {
		green = color.New(color.FgGreen).SprintFunc()
		fmt.Print(green("[+] "))
		fmt.Println(v...)
	} else {
		fmt.Print("[+] ")
		for _, arg := range v {
			fmt.Print(arg)
		}
		fmt.Print("\n")
	}
}

func Badln(v ...interface{}) {
	if runtime.GOOS == "linux" {
		red = color.New(color.FgRed).SprintFunc()
		fmt.Print(red("[-] "))
		fmt.Println(v...)
	} else {
		fmt.Print("[-] ")
		fmt.Println(v...)
	}
}

func Infoln(v ...interface{}) {
	if runtime.GOOS == "linux" {
		purple = color.New(color.FgMagenta).SprintFunc()
		fmt.Print(purple("[*] "))
		fmt.Println(v...)
	} else {
		fmt.Print("[*] ")
		fmt.Println(v...)
	}
}

func Greenln(v ...interface{}) {
	green = color.New(color.FgGreen).SprintFunc()
	fmt.Println(green(v...))
}

func Redln(v ...interface{}) {
	red = color.New(color.FgRed).SprintFunc()
	fmt.Println(red(v...))
}

func Purpleln(v ...interface{}) {
	purple = color.New(color.FgMagenta).SprintFunc()
	fmt.Println(purple(v...))
}

func Blueln(v ...interface{}) {
	blue = color.New(color.FgBlue).SprintFunc()
	fmt.Println(blue(v...))
}

func Cyanln(v ...interface{}) {
	cyan = color.New(color.FgCyan).SprintFunc()
	fmt.Println(cyan(v...))
}

// == Separator between Println and Print functions ==

func Good(v ...interface{}) {
	if runtime.GOOS == "linux" {
		green = color.New(color.FgCyan).SprintFunc()
		fmt.Print(green("[+] "))
		fmt.Println(v...)
	} else {
		fmt.Print("[+] ")
		fmt.Println(v...)
	}
}

func Bad(v ...interface{}) {
	if runtime.GOOS == "linux" {
		red = color.New(color.FgRed).SprintFunc()
		fmt.Print(red("[-] "))
		fmt.Println(v...)
	} else {
		fmt.Print("[-] ")
		fmt.Println(v...)
	}
}

func Info(v ...interface{}) {
	if runtime.GOOS == "linux" {
		purple = color.New(color.FgMagenta).SprintFunc()
		fmt.Print(purple("[*] "))
		fmt.Println(v...)
	} else {
		fmt.Print("[*] ")
		fmt.Println(v...)
	}
}

func Green(v ...interface{}) {
	green = color.New(color.FgGreen).SprintFunc()
	fmt.Print(green(v...))
}

func Red(v ...interface{}) {
	red = color.New(color.FgRed).SprintFunc()
	fmt.Print(red(v...))
}

func Purple(v ...interface{}) {
	purple = color.New(color.FgMagenta).SprintFunc()
	fmt.Print(purple(v...))
}

func Blue(v ...interface{}) {
	blue = color.New(color.FgBlue).SprintFunc()
	fmt.Print(blue(v...))
}

func Cyan(v ...interface{}) {
	cyan = color.New(color.FgCyan).SprintFunc()
	fmt.Print(cyan(v...))
}

// == This functions return the coloured input ==

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
