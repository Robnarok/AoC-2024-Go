package helper

import (
	"bufio"
	"log"
	"os"
)

func Parse(filename string) []string {
	file, err := os.Open("../input/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var content []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content
}
