package main

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {}

// GOOD
func useHTTPS() {
	remoteURL := "https://example.com/hello.bin"
	localFilepath := "./hello.bin"
	err := downloadFile(localFilepath, remoteURL)
	if err != nil {
		panic(err)
	}

	command := exec.Command("bash", localFilepath)
	_ = command
}

// BAD
func useHTTP() {
	remoteURL := "http://example.com/hello.bin"
	localFilepath := "./hello.bin"
	err := downloadFile(localFilepath, remoteURL)
	if err != nil {
		panic(err)
	}

	command := exec.Command("bash", localFilepath)
	_ = command
}

// GOOD
func useFuncGood() {
	utilDownloadAndExecute("https://example.com/hello.bin")
}

// BAD
func useFuncBad() {
	utilDownloadAndExecute("http://example.com/hello.bin")
}

// BAD
func downloadAndExecuteWithPrefixBad() {
	utilDownloadAndExecuteWithPrefix("./hello.bin")
}

// GOOD
func noExecution() {
	remoteURL := "http://example.com/hello.bin"
	localFilepath := "./hello.bin"
	err := downloadFile(localFilepath, remoteURL)
	if err != nil {
		panic(err)
	}
}

func utilDownloadAndExecute(remoteURL string) {
	localFilepath := "./hello.bin"
	err := downloadFile(localFilepath, remoteURL)
	if err != nil {
		panic(err)
	}

	command := exec.Command("bash", localFilepath)
	_ = command
}
func utilDownloadAndExecuteWithPrefix(name string) {
	remoteURL := "http://example.com/" + name
	localFilepath := "./hello.bin"
	err := downloadFile(localFilepath, remoteURL)
	if err != nil {
		panic(err)
	}

	command := exec.Command("bash", localFilepath)
	_ = command
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
