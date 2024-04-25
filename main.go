package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var byteCount, lineCount, wordCount bool

func init() {
	flag.BoolVar(&byteCount, "c", false, "get character count")
	flag.BoolVar(&lineCount, "l", false, "get character count")
	flag.BoolVar(&wordCount, "w", false, "get character count")
}

func main() {
	flag.Parse()
	fileName := flag.Arg(0)

	if byteCount {
		count, err := getByteCount(fileName)
		if err != nil {
			fmt.Println(err, fileName)
			return
		}

		fmt.Println(count, fileName)
	}

	if lineCount {
		count, err := getLineCount(fileName)
		if err != nil {
			fmt.Println(err, fileName)
			return
		}

		fmt.Println(count, fileName)
	}

	if wordCount {
		count, err := getWordCount(fileName)
		if err != nil {
			fmt.Println(err, fileName)
			return
		}

		fmt.Println(count, fileName)
	}
}

func getByteCount(fileName string) (int, error) {
	info, err := os.Stat(fileName)
	if err != nil {
		return -1, err
	}

	return int(info.Size()), nil
}

func getLineCount(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return -1, err
	}

	defer file.Close()

	var count int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count += 1
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return count, nil
}

func getWordCount(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return -1, err
	}

	defer file.Close()

	var count int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		count += len(words)
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return count, nil
}
