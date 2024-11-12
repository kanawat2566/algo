package wordcount

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFileText() []string {
	// file, err := os.Open("D:/workTraning-GO/AlgorithmDesign/WordCount/test.txt")
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current working directory:", dir)

	file, err := os.Open("WordCount/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	return strings.Split(string(content), " ")
}
