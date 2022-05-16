package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Running %s go on %s\n", os.Args[0], os.Getenv("GOFILE"))
}
