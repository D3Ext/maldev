package main

import (
	"strconv"
	"strings"
	"time"
  "log"
  "fmt"

	l "github.com/D3Ext/maldev/src/logging"
	"github.com/D3Ext/maldev/src/redteam"
)

func main() {
	var cfg redteam.SandboxConfig = redteam.DefaultConfig // Define default config to use
	var scanning_results redteam.SandboxResults        // Define sandboxing results

	var checks_counter int // A counter is declared to sum all checks
	var disk_space int
	var mem int
	var dt int
	var drivers_check bool
	var proc_check bool
	var driver_name string
	var proc_name string
	var err error

	l.PrintBanner("go-sandbox") // Directly print banner format
	time.Sleep(200 * time.Millisecond)
	fmt.Println()
	l.Info("Starting sandbox analysis") // Status bar
	time.Sleep(400 * time.Millisecond)

	l.Info("Getting disk maximum space...")
	if cfg.MinSpace != 0 && cfg.Path != "" { // Check disk usage
		disk_space, err = redteam.DiskCheck(cfg.Path)
		if err != nil {
			log.Fatal(err)
		}

		if disk_space < cfg.MinSpace { // Check if system space is lower than especified one
			scanning_results.MinSpace = true
		} else {
			scanning_results.MinSpace = false
		}
	}
	time.Sleep(300 * time.Millisecond)

	l.Info("Retrieving CPU cores...")
	if cfg.Cores != 0 { // Check number of CPU cores and compare
		if redteam.CpuCheck() < cfg.Cores {
			scanning_results.Cores = true
			checks_counter = checks_counter + 1
		} else {
			scanning_results.Cores = false
		}
	}
	time.Sleep(300 * time.Millisecond)

	l.Info("Checking total memory usage...")
	if cfg.Memory != 0 { // Get total memory usage
		mem, err = redteam.MemoryCheck()
		if err != nil {
			log.Fatal(err)
		}

		if mem < cfg.Memory {
			scanning_results.Memory = true
			checks_counter = checks_counter + 1
		} else {
			scanning_results.Memory = false
		}
	}
	time.Sleep(300 * time.Millisecond)

	l.Info("Looking for possible \"Sleep Skipping\" techniques...")
	if cfg.Time != 0 { // Check time (detect Sleep Skipping technique)
		dt = redteam.TimeCheck(cfg.Time)
		if dt-50 > cfg.Time/2 {
			scanning_results.Time = false
		} else {
			scanning_results.Time = true
			checks_counter = checks_counter + 1
		}
	}

	l.Info("Looking for potential drivers (Virtual Box, VMWare, Qemu)")
	if cfg.Drivers == true { // Look for common drivers (Virtual Box, VMWare, Qemu...)
		drivers_check, driver_name = redteam.DriversCheck()
		if drivers_check == true {
			scanning_results.Drivers = true
			checks_counter = checks_counter + 1
		} else {
			scanning_results.Drivers = false
		}
	}
	time.Sleep(300 * time.Millisecond)

	l.Info("Checking internet connection...")
	if cfg.Internet == true { // Check if there is internet connection
		conn_check := redteam.InternetCheck()
		if conn_check == false {
			scanning_results.Internet = true
			checks_counter = checks_counter + 1
		} else {
			scanning_results.Internet = false
		}
	}
	time.Sleep(300 * time.Millisecond)

	l.Info("Retrieving all processes to check if any virtualization app is running")
	if cfg.Processes == true { // Look for potentially analysis apps processes
		proc_check, proc_name, err = redteam.ProcessesCheck()
		if err != nil {
			log.Fatal(err)
		}

		if proc_check == true {
			scanning_results.Processes = true
			checks_counter = checks_counter + 1
		} else {
			scanning_results.Processes = false
		}
	}
	time.Sleep(400 * time.Millisecond)

	l.Success("Sandbox analysis completed!")
	fmt.Println()
	time.Sleep(200 * time.Millisecond)

	// Start displaying results
	if scanning_results.MinSpace == true {
		fmt.Println("Disk usage is lower than expected! It's probably a sandbox: \t-")
	} else {
		disk_gb := disk_space / 1024 / 1024 / 1024
		fmt.Println("Disk usage seems to be perfect (", disk_gb, "GB ): \t+")
	}
	time.Sleep(100 * time.Millisecond)

	if scanning_results.Cores == true {
		fmt.Println("System has", redteam.CpuCheck(), "processors, it may be a sandbox: \t-")
	} else {
		fmt.Println("System has", redteam.CpuCheck(), "processors, it seems legit: \t+")
	}
	time.Sleep(100 * time.Millisecond)

	if scanning_results.Memory == true {
		fmt.Println("Memory is lower than expected! May be a sandbox: \t-")
	} else {
		fmt.Println(mem, "bytes of memory, it seems legit: \t+")
	}

	if scanning_results.Processes == true {
		fmt.Println("Dangerous process was found (", proc_name, "): \t-")
	} else {
		fmt.Println("No dangerous process was found: \t\t+")
	}

	if scanning_results.Time == true {
		fmt.Println("An sleeping issue was detected, sleeping time was lower than in config: \t-")
	} else {
		fmt.Println("Sleeping results are OK (", dt, "ms ): \t\t+")
	}

	if scanning_results.Drivers == true {
		driver_name = strings.Split(driver_name, "\\")[len(strings.Split(driver_name, "\\"))-1] // Get filename of the driver
		l.Println("Virtualization driver was found (", driver_name, "): \t-")
	} else {
		l.Println("No virtualization driver was found so system should be real: \t+")
	}

	if scanning_results.Internet == true {
		fmt.Println("System doesn't have internet connection: \t\t-")
	} else {
		fmt.Println("System has internet connection: \t\t+")
	}

	fmt.Println()
	l.Info("Sandboxing triggers found: " + strconv.Itoa(checks_counter) + "/7")
}
