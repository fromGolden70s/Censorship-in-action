package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var fileName string
	fmt.Scanln(&fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	words := make(map[string]struct{})

	for scanner.Scan() {
		words[strings.ToLower(scanner.Text())] = struct{}{}
	}
	file.Close()

	checkWords(words)
}

func checkWords(words map[string]struct{}) {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	splitInput := strings.Split(input, " ")

	s := ""

	if strings.EqualFold(strings.TrimSpace(splitInput[0]), "exit") {
		fmt.Print("Bye!")

	} else {
		for _, val := range splitInput {

			if _, ok := words[strings.TrimSpace(strings.ToLower(val))]; ok {

				for i := 1; i <= len(strings.TrimSpace(val)); i++ {
					s = s + "*"
				}
				s += " "

			} else {
				s = s + val + " "
			}
		}
		strings.TrimSpace(s)
		fmt.Println(s)

		checkWords(words)
	}
}
