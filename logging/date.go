package logging

import (
  "time"
  "fmt"
)

func GetTime() (string) {
  return time.Now().Format("15:04:05")
}

func TimePrintln(input string) {
  current_date := time.Now().Format("15:04:05")
  fmt.Println(current_date + " " + input)
}

func TimePrint(input string) {
  current_date := time.Now().Format("15:04:05")
  fmt.Print(current_date + " " + input)
}

func TimeSprint(input string) (string) {
  current_date := time.Now().Format("15:04:05")
  return fmt.Sprint(current_date + " " + input)
}


