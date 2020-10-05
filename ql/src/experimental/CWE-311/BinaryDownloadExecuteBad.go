package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {}

func bad() {
	remoteURL := "http://example.com/hello.bin"
	localFilepath := "./hello.bin"
	err := downloadFile(localFilepath, remoteURL)
	if err != nil {
		panic(err)
	}
	command := exec.Command("bash", localFilepath)
	_ = command
	//...
}

func downloadFile(filepath string, url string) error {
	// Send request:
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create a new file:
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write response body to file:
	_, err = io.Copy(out, resp.Body)
	return err
}
