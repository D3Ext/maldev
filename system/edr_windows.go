package system

/*

References:
https://github.com/FourCoreLabs/EDRHunt

*/

import (
  "github.com/FourCoreLabs/EDRHunt/pkg/edrRecon"
  "github.com/FourCoreLabs/EDRHunt/pkg/resources"
  "strings"
  "fmt"
)

type EdrInfo struct {
  Processes   []resources.ProcessMetaData

/*type ProcessMetaData struct {
	  ProcessName        string
	  ProcessPath        string
	  ProcessDescription string
	  ProcessCaption     string
  	ProcessCmdLine     string
	  ProcessPID         string
	  ProcessParentPID   string
	  ProcessExeMetaData FileMetaData
    ScanMatch          []string
  }*/

  Drivers     []resources.DriverMetaData
  
/*type DriverMetaData struct {
	  DriverBaseName    string
	  DriverSysMetaData FileMetaData
	  DriverFilePath    string
	  ScanMatch         []string
  }*/

  Services    []resources.ServiceMetaData

/*type ServiceMetaData struct {
	  ServiceName        string
	  ServiceDisplayName string
	  ServiceDescription string
	  ServiceCaption     string
	  ServicePathName    string
	  ServiceState       string
	  ServiceProcessId   string
	  ServiceExeMetaData FileMetaData
	  ScanMatch          []string
  }*/

}

func GetEdrInfo() (*EdrInfo, error) {
  processes_info, err := edrRecon.CheckProcesses()
  if err != nil {
    return new(EdrInfo), err
  }

  services_info, err := edrRecon.CheckServices()
  if err != nil {
    return new(EdrInfo), err
  }

  drivers_info, err := edrRecon.CheckDrivers()
  if err != nil {
    return new(EdrInfo), err
  }

  return &EdrInfo{
    Processes: processes_info,
    Drivers: drivers_info,
    Services: services_info,
  }, nil
}

func (e *EdrInfo) Format() string {
  var out string
  for _, proc := range e.Processes {
    out += fmt.Sprintf("Suspicious Process Name: %s\n", proc.ProcessName)
    out += fmt.Sprintf("Description: %s\n", proc.ProcessDescription)
    out += fmt.Sprintf("Caption: %s\n", proc.ProcessCaption)
    out += fmt.Sprintf("Binary: %s\n", proc.ProcessPath)
    out += fmt.Sprintf("ProcessID: %s\n", proc.ProcessPID)
    out += fmt.Sprintf("Parent Process: %s\n", proc.ProcessParentPID)
    out += fmt.Sprintf("Process CmdLine: %s\n\n", proc.ProcessCmdLine)
  }

  for _, driver := range e.Drivers {
    out += fmt.Sprintf("Suspicious Driver Module: %s\n", driver.DriverBaseName)
    out += fmt.Sprintf("Driver FilePath: %s\n", driver.DriverFilePath)
    out += fmt.Sprintf("Driver File Metadata: \t%s\n", edrRecon.FileMetaDataParser(driver.DriverSysMetaData))
    out += fmt.Sprintf("Matched Keyword: %s\n\n", driver.ScanMatch)
  }

  for _, service := range e.Services {
    out += fmt.Sprintf("Suspicious Service Name: %s\n", service.ServiceName)
    out += fmt.Sprintf("Display Name: %s\n", service.ServiceDisplayName)
    out += fmt.Sprintf("Caption: %s\n", service.ServiceCaption)
    out += fmt.Sprintf("CommandLine: %s\n", service.ServicePathName)
    out += fmt.Sprintf("Status: %s\n", service.ServiceState)
    out += fmt.Sprintf("ProcessID: %s\n\n", service.ServiceProcessId)
  }

  out = strings.TrimSuffix(out, "\n\n")
  return out
}

