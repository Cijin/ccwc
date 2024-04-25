package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

var byteCount, lineCount, wordCount, charCount bool

func init() {
	flag.BoolVar(&byteCount, "c", false, "get number of bytes in file")
	flag.BoolVar(&lineCount, "l", false, "get line count")
	flag.BoolVar(&wordCount, "w", false, "get word count")
	flag.BoolVar(&charCount, "m", false, "get character count")
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

	if charCount {
		count, err := getCharCount(fileName)
		if err != nil {
			fmt.Println(err, fileName)
			return
		}

		fmt.Println(count, fileName)
	}
}

func getByteCount(fileName string) (int, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return -1, err
	}

	return len(data), nil
}

func getLineCount(fileName string) (int, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return -1, err
	}

	lines := strings.Split(string(data), "\n")

	// ignore the string after the last line split (empty), hence -1
	return len(lines) - 1, nil
}

func getWordCount(fileName string) (int, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return -1, err
	}

	fields := strings.Fields(string(data))

	return len(fields), nil
}

func getCharCount(fileName string) (int, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return -1, err
	}

	return utf8.RuneCount(data), nil
}
