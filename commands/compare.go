package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Difference(a, b []byte) (diff []byte) {
	m := make(map[byte]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}

func main() {
	const colorRed string = "\033[31m"

	fileData1, err := ioutil.ReadFile("./1.txt")
	check(err)
	fileData2, err := ioutil.ReadFile("./2.txt")
	check(err)

	if string(Difference(fileData1, fileData2)) == "" {
		fmt.Println("Files are the same!")
	} else {
		fmt.Println(string(colorRed), "First file: \n", string(Difference(fileData1, fileData2)))
		fmt.Println(string(colorRed), "Second file: \n", string(Difference(fileData2, fileData1)))
	}
}
