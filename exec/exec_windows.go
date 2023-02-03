package exec

import (
  ex "os/exec"
  "syscall"

  "golang.org/x/sys/windows"
)

func ExecuteCommand(comm string) (string) { // Custom function to execute commands
  var out []byte
  cmd := ex.Command("cmd.exe", "/c", comm) // Create command
  cmd.SysProcAttr = &syscall.SysProcAttr{
    HideWindow: true,
  } // Set attributes, hidden window

  out, _ = cmd.CombinedOutput() // Get output (stderr and stdout)
  return string(out)
}

func ExecutePowershell(comm string) (string) { // Custom function to execute commands
  var out []byte
  cmd := ex.Command("powershell.exe", comm) // Create command
  cmd.SysProcAttr = &syscall.SysProcAttr{
    HideWindow: true,
  } // Set attributes, hidden window

  out, _ = cmd.CombinedOutput() // Get output (stderr and stdout)
  return string(out)
}

func ExecWithToken(comm string, token windows.Token) (string) {
  var out []byte
  cmd := ex.Command("powershell.exe", comm)
  cmd.SysProcAttr = &syscall.SysProcAttr{
    HideWindow: true,
    Token: syscall.Token(token),
  }

  out, _ = cmd.CombinedOutput()
  return string(out)
}
