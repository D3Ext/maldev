package exec

/*

References:
https://stackoverflow.com/questions/1877045/how-do-you-get-the-output-of-a-system-command-in-go
https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html

*/

import (
  ex "os/exec"
)

func ExecuteCommand(comm string) (string) { // Return stdout and stderr
  var out []byte
  cmd := ex.Command("/bin/bash", "-c", comm)
  out, _ = cmd.CombinedOutput()
  return string(out)
}


