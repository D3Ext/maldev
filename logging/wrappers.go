package logging

/*

This file provides same logging functions as "fmt" and "log" native packages

References:
https://pkg.go.dev/log
https://pkg.go.dev/fmt

*/

import (
	"fmt"
	"io"
	"log"
)

// fmt functions

func Print(v ...any) {
	fmt.Print(v...)
}

func Println(v ...any) {
	fmt.Println(v...)
}

func Printf(format string, v ...any) (int, error) {
	return fmt.Printf(format, v...)
}

func Sprintf(format string, v ...any) string {
	return fmt.Sprintf(format, v...)
}

func Fprint(w io.Writer, a ...any) (n int, err error) {
	return fmt.Fprint(w, a...)
}

func Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
}

func Append(b []byte, a ...any) []byte {
	return fmt.Append(b, a...)
}

func Appendf(b []byte, format string, a ...any) []byte {
	return fmt.Appendf(b, format, a...)
}

func Appendln(b []byte, a ...any) []byte {
	return fmt.Appendln(b, a...)
}

// log functions

func Panic(v ...any) {
	log.Panic(v...)
}

func Panicf(format string, v ...any) {
	log.Panicf(format, v...)
}

func Panicln(v ...any) {
	log.Panicln(v...)
}

func Fatal(v ...any) {
	log.Fatal(v...)
}

func Fatalln(v ...any) {
	log.Fatalln(v...)
}

func Fatalf(format string, v ...any) {
	log.Fatalf(format, v...)
}
