package system

/*

References:
https://github.com/Ne0nd0g/merlin-agent/blob/dev/commands/uptime_windows.go
https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/nf-sysinfoapi-gettickcount64

*/

import (
  "strconv"
  "strings"
  "time"

  "golang.org/x/sys/windows"
)

type Upt struct {
  Duration    time.Duration
}

func Uptime() (int) {
  kernel32 := windows.NewLazySystemDLL("kernel32")
  GetTicketCount64 := kernel32.NewProc("GetTickCount64")

  r1, _, _ := GetTicketCount64.Call(0, 0, 0, 0)
	
  uptime_time := time.Duration(r1) * time.Millisecond
  uptime_int, _ := strconv.Atoi(uptime_time.String())

  return uptime_int
}

func (u Upt) Format() (string) {
  fr := strings.Replace(u.Duration.String(), "h", "h ", 1)
  fr = strings.Replace(u.Duration.String(), "m", "m ", 1)

  return fr
}

