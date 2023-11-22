package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	shell := flag.String("shell", "bash", "The shell you want to syncronize")

	flag.Parse()

	fmt.Println("Using", *shell)

	path := "/usr/share/syncrc/"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	fmt.Printf("from: /home/marcus/.%src\n", *shell)
	fmt.Printf("to: %s.%src.bkp\n", path, *shell)

	nBytes, err := copy(fmt.Sprintf("/home/marcus/.%src", *shell), fmt.Sprintf("%s.%src.bkp", path, *shell))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(nBytes, "copied.")
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
