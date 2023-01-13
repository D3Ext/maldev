package system

import (
  //"fmt"
  "os/user"
  "golang.org/x/sys/windows"
  sysinfo "github.com/elastic/go-sysinfo"
  "github.com/fourcorelabs/wintoken"
)

type UserInfo struct { // Struct for WhoamiAll function
  Username            string
  IntegrityLevel      string
  Privileges          []wintoken.Privilege
}

type MemInfo struct {
  Total     uint64
  Used      uint64
  Available uint64
}

type SystemInfo struct {
  Hostname      string
  OsName        string
  OsVersion     string
  Architecture  string
  KernelVersion string
  Timezone      string
  ComputerId    string
  Container     bool
  Memory        MemInfo
}

func GetUserPrivs() (bool, error) { // Function to check if current user has Administrator privileges
  var sid *windows.SID
  err := windows.AllocateAndInitializeSid(
    &windows.SECURITY_NT_AUTHORITY,
    2,
    windows.SECURITY_BUILTIN_DOMAIN_RID,
    windows.DOMAIN_ALIAS_RID_ADMINS,
    0, 0, 0, 0, 0, 0,
    &sid)

  if err != nil {
    return false, err
  }

  token := windows.Token(0)
  member, err := token.IsMember(sid) // Check if is inside admin group
  if err != nil {
    return false, err
  }
  return member, nil // return true or false
}

func Systeminfo() (*SystemInfo, error) {
  //var final_info string
  host, err := sysinfo.Host()
  if err != nil {
    return new(SystemInfo), err
  }

  var container_bool bool
  if host.Info().Containerized == nil {
    container_bool = false
  } else {
    container_bool = true
  }

  mem, err := host.Memory()
  if err != nil {
    return new(SystemInfo), err
  }

  memory_info := MemInfo{
    Total: mem.Total,
    Used: mem.Used,
    Available: mem.Available,
  }

  return &SystemInfo{
    Hostname: host.Info().Hostname,
    OsName: host.Info().OS.Name,
    OsVersion: host.Info().OS.Version,
    Architecture: host.Info().Architecture,
    KernelVersion: host.Info().KernelVersion,
    Timezone: host.Info().Timezone,
    ComputerId: host.Info().UniqueID,
    Container: container_bool,
    Memory: memory_info,
  }, nil
}

func WhoamiAll() (*UserInfo, error) {
  token, err := wintoken.OpenProcessToken(0, wintoken.TokenPrimary)
  if err != nil {
    return new(UserInfo), err
  }
  defer token.Close()

  // Get integrity level
  integrity_level, err := token.GetIntegrityLevel()
  if err != nil {
    return new(UserInfo), err
  }

  // Get privileges
  privs, err := token.GetPrivileges()
  if err != nil {
    return new(UserInfo), err
  }

  localuser, err := user.Current()
  if err != nil {
    return new(UserInfo), err
  }

  return &UserInfo{
    Username: localuser.Username,
    IntegrityLevel: integrity_level,
    Privileges: privs,
  }, nil

}

