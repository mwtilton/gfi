package main

import (
	"fmt"
	"golang.org/x/mod/sumdb/dirhash"
)

func main() {
	fmt.Println("Hello, world.")
	
}

func HashZip(zipfile string, hash Hash) (string, error)