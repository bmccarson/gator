package main

import (
	"fmt"

	"github.com/bmccarson/gator/internal/config"
)

func main() {
	file, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	file.SetUser("Brian")

	newFile, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(newFile)
}
