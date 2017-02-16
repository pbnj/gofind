package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

var (
	startPath string
	filename  string
)

func init() {
	usr, _ := user.Current()
	usrHomeDir := usr.HomeDir
	flag.StringVar(&startPath, "p", usrHomeDir, "Starting Path")
	flag.StringVar(&filename, "n", "", "Filename")
}

func main() {
	flag.Parse()
	if filename == "" {
		fmt.Println("Filename cannot be blank.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	filepath.Walk(startPath, walkFunc)
}

func walkFunc(path string, f os.FileInfo, err error) error {
	data, err := ioutil.ReadFile(path + "/" + filename)
	if err != nil {
		// skip level
	} else {
		fmt.Println("Found .mysql_history in:", path)
		fmt.Println(string(data))
	}

	return nil
}
