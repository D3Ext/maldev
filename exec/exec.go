package exec

import (
  ex "os/exec"
)

func ExecuteCommand(comm string) (string) { // Return stdout and stderr
  var out []byte
  cmd := ex.Command("/bin/bash", "-c", comm)
  out, _ = cmd.CombinedOutput()
  return string(out)
}


