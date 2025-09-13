package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bmccarson/gator/internal/commands"
	"github.com/bmccarson/gator/internal/config"
	"github.com/bmccarson/gator/internal/state"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	data := state.Init(cfg)
	inputCommands := commands.Init()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Gator> ")

		scanner.Scan()
		input := cleanInput(scanner.Text())

		command := input[0]
		arg := ""
		if len(input) == 2 {
			arg = input[1]
		}

		if key, exists := inputCommands[command]; exists {
			err := key.Callback(&data, arg)

			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("command does not exisist")
		}
	}
}

func cleanInput(text string) []string {
	cleanedInput := []string{}

	words := strings.Fields(text)

	for _, word := range words {
		clean := strings.ToLower(word)

		cleanedInput = append(cleanedInput, clean)
	}

	return cleanedInput
}
