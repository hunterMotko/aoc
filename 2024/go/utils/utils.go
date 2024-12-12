package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
  defer file.Close()
	scn := bufio.NewScanner(file)
	var lines []string
	for scn.Scan() {
		lines = append(lines, scn.Text())
	}
	return lines, nil
}

func ReadFileToMatrix(name string) ([][]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
  defer file.Close()
	scn := bufio.NewScanner(file)
	var matrix [][]string
	for scn.Scan() {
		matrix = append(matrix, strings.Split(scn.Text(), ""))
	}
	return matrix, nil
}

func ReadFileToIntMatrix(name string) ([][]int, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
  defer file.Close()
	scn := bufio.NewScanner(file)
	var matrix [][]int
	for scn.Scan() {
    var temp []int
    for _, b := range scn.Bytes() {
      temp = append(temp, int(b - '0'))
    }
    matrix = append(matrix, temp)
	}
	return matrix, nil
}

func ParseInt(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 0)
}

func StrToint(line string, del string) []int {
  var res []int
  arr := strings.Split(line, del)
  for _, v := range arr {
    n, err := strconv.Atoi(v)
    if err != nil {
      log.Fatalf("PARSE ERROR: %v\n", err)
    }
    res = append(res, n)
  }
  return res
}
