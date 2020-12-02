package utils

import (
	"io/ioutil"
	"strconv"
	strings "strings"
)

func fileToSlice(file string) ([]string, error) {

	res := []string{}
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return res, err
	}

	data := strings.Split(string(dat), "\r\n")
	return data, err
}

func sliceToInt(stringSlice []string) []int {

	res := []int{}
	for _, element := range stringSlice {

		value, _ := strconv.Atoi(element)
		res = append(res, value)
	}
	return res
}

func ReadFileAsInts(file string) ([]int, error) {

	stringSlice, err := fileToSlice("input/day1.txt")

	return sliceToInt(stringSlice), err
}
