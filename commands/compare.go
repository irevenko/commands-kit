package main

import (
	"fmt"
	"io/ioutil"
)

func readFromFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func main() {
	readFromFile("./1.txt")
	readFromFile("./2.txt")
}
