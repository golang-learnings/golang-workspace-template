package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var notes []string
loop:
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter command and data: ")

		input, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}

		line := string(input)
		arr := strings.Split(line, " ")

		var initialX int

		if strings.Contains(line, ">") {
			initialX = 1
		} else {
			initialX = 0
		}

		command := arr[initialX]
		note := ""

		for i := initialX + 1; i < len(arr); i++ {
			if i != len(arr)-1 {
				note += arr[i] + " "
			} else {
				note += arr[i]
			}
		}

		// fmt.Print("Command "+command, ",", " Note "+note)

		if command == "create" {
			notes = append(notes, note)
			fmt.Println("create")
		} else if command == "list" || strings.Contains(line, "list") {
			for i := 0; i < len(notes); i++ {
				fmt.Print("NOTE " + string(i) + " " + notes[i])
			}
		} else if command == "exit" || strings.Contains(line, "exit") {
			fmt.Print("[Info] Bye!")
			break loop
		} else {
			fmt.Println("Enter a Valid Command ( list, create, exit )")
		}
	}
}
