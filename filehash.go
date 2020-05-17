package main

import (
	"fmt"
	"golang.org/x/mod/sumdb/dirhash"
)

func main() {

	filePath := flag.String("directory", "/path", "file path location")

	//numbPtr := flag.Int("numb", 42, "an int")
	//boolPtr := flag.Bool("fork", false, "a bool")

	//var svar string
	//flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("tail:", flag.Args())
}
	