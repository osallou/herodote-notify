package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "herodote notify client"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func dirExists(name string) bool {
	if fi, err := os.Stat(name); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}
	return false
}

var Version string

func main() {
	var server string
	var helpVersion = false

	flag.BoolVar(&helpVersion, "version", false, "Show version")
	flag.StringVar(&server, "url", "", "herodote job status url")
	flag.Parse()

	if helpVersion {
		fmt.Printf("Version: %s\n", Version)
		return
	}

	byteData := make([]byte, 0)
	req, _ := http.NewRequest("PUT", server, bytes.NewReader(byteData))
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to contact server %s: %s\n", server, err)
	}

}
