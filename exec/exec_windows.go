package exec

import (
  ex "os/exec"
  "syscall"
)

func ExecutePowershell(comm string) string { // Custom function to execute commands
  var out []byte
  cmd := ex.Command("powershell.exe", comm) // Create command
  cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true,} // Set attributes, hidden window

  out, _ = cmd.CombinedOutput() // Get output (stderr and stdout)
  return string(out)
}

