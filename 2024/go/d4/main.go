package main

import (
	"aoc/utils"
	"fmt"
	"log"
)

func check(matrix [][]string, i, j int) int {
	count := 0
	// horizontal backwards
	if j >= 3 &&
		matrix[i][j-1] == "M" &&
		matrix[i][j-2] == "A" &&
		matrix[i][j-3] == "S" {
		count++
	}
	// horizontal forward
	if j <= (len(matrix[0])-1-3) &&
		matrix[i][j+1] == "M" &&
		matrix[i][j+2] == "A" &&
		matrix[i][j+3] == "S" {
		count++
	}
	// vertical upwards
	if i >= 3 &&
		matrix[i-1][j] == "M" &&
		matrix[i-2][j] == "A" &&
		matrix[i-3][j] == "S" {
		count++
	}
	// vertical downwards
	if i <= (len(matrix[0])-1-3) &&
		matrix[i+1][j] == "M" &&
		matrix[i+2][j] == "A" &&
		matrix[i+3][j] == "S" {
		count++
	}
	// diagonal upwards right
	if i >= 3 && j <= (len(matrix[0])-1-3) &&
		matrix[i-1][j+1] == "M" &&
		matrix[i-2][j+2] == "A" &&
		matrix[i-3][j+3] == "S" {
		count++
	}
	// diagonal upwards left
	if i >= 3 && j >= 3 &&
		matrix[i-1][j-1] == "M" &&
		matrix[i-2][j-2] == "A" &&
		matrix[i-3][j-3] == "S" {
		count++
	}
	// diagonal downwards right
	if i <= (len(matrix[0])-1-3) && j <= (len(matrix[0])-1-3) &&
		matrix[i+1][j+1] == "M" &&
		matrix[i+2][j+2] == "A" &&
		matrix[i+3][j+3] == "S" {
		count++
	}
	// diagonal downwards left
	if i <= (len(matrix[0])-1-3) && j >= 3 &&
		matrix[i+1][j-1] == "M" &&
		matrix[i+2][j-2] == "A" &&
		matrix[i+3][j-3] == "S" {
		count++
	}
	return count
}
func part1(matrix [][]string) int {
	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "X" {
				res += check(matrix, i, j)
			}
		}
	}
	return res
}

func checkX(matrix [][]string, i, j int) int {
	res := 0
	if j < 1 || j > len(matrix[i])-2 {
		return 0
	}
	if matrix[i-1][j+1] == "M" &&
		matrix[i-1][j-1] == "M" &&
		matrix[i+1][j-1] == "S" &&
		matrix[i+1][j+1] == "S" {
		res++
	}
	if matrix[i-1][j+1] == "S" &&
		matrix[i-1][j-1] == "S" &&
		matrix[i+1][j-1] == "M" &&
		matrix[i+1][j+1] == "M" {
		res++
	}
	if matrix[i-1][j+1] == "S" &&
		matrix[i-1][j-1] == "M" &&
		matrix[i+1][j-1] == "M" &&
		matrix[i+1][j+1] == "S" {
		res++
	}
	if matrix[i-1][j+1] == "M" &&
		matrix[i-1][j-1] == "S" &&
		matrix[i+1][j-1] == "S" &&
		matrix[i+1][j+1] == "M" {
		res++
	}
	return res
}
func part2(matrix [][]string) int {
	res := 0
	for i := 1; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "A" {
				res += checkX(matrix, i, j)
			}
		}
	}
	return res
}

func main() {
	matrix, err := utils.ReadFileToMatrix("./d4/in.txt")
	if err != nil {
		log.Fatalf("READ ERROR: %v\n", err)
	}
	// fmt.Println(part1(matrix))
	fmt.Println(part2(matrix))
}
