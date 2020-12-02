package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	strings "strings"
)

func sliceToInt(stringSlice []string) (asIntegers []int) {

	for _, element := range stringSlice {
		asIntegers = append(asIntegers, AtoiPanic(element))
	}
	return
}

func FileAsLines(file string) []string {

	stream, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Cannot read file:", file)
		panic(err)
	}

	return strings.Split(string(stream), "\r\n")
}

func FileAsInts(file string) []int {

	stringSlice := FileAsLines(file)
	return sliceToInt(stringSlice)
}

func ResultStrings(value1, value2 int) (string, string) {
	return strconv.Itoa(value1), strconv.Itoa(value2)
}

func AtoiPanic(str string) int {

	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Cannot convert to int", str)
		panic(err)
	}

	return i
}
