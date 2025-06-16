package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
)

func main() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)

		}

		file, err := os.Create("Keystrokes.txt")
		if err != nil {
			fmt.Println(err)
		}
		_, err = file.WriteString("\n" + string(char))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println(string(key))

	}

}
