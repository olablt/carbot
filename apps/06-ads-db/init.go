package main

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func setupPath(path string) string {
	if err := initDir(path); err != nil {
		log.Fatal(err)
	}
	return path
}
