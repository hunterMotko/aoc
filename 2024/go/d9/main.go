package main

import (
	_ "embed"
	"fmt"
)

//go:embed in.txt
var test string

func main() {
	fmt.Println(part1(test)) // 6 262 891 638 328
	fmt.Println(part2(test))
}

func orderDisk(in string) []int {
	var disks []int
	j := 0
	for i, v := range in {
		n := int(v - '0')
		if n < 0 {
			break
		}
		if i == 0 {
			for ; n > 0; n-- {
				disks = append(disks, j)
			}
			j++
		} else if i%2 == 1 {
			for ; n > 0; n-- {
				disks = append(disks, -1)
			}
		} else {
			for ; n > 0; n-- {
				disks = append(disks, j)
			}
			j++
		}
	}
	return disks
}

func part1(in string) int {
	temp := orderDisk(in)
	j := len(temp) - 1
	res := 0
	for i := 0; i < len(temp); i++ {
		if i == j {
			break
		}
		if temp[i] == -1 {
			for ; j >= 0; j-- {
				if temp[j] != -1 {
					c := temp[j]
					temp[j] = temp[i]
					temp[i] = c
					break
				}
			}
		}
		if temp[i] != -1 {
			res += i * temp[i]
		}
	}
	return res
}

type file struct {
	i, len int
}

func parse(in string) ([]int, []file, []file) {
	diskMap := []byte(in)
	isFile := true
	var fileID int
	var decoded []int
	var files []file
	var spaces []file
	for _, v := range diskMap[:len(diskMap)-1] {
		n := int(v - '0')
		if isFile {
			isFile = false
			files = append(files, file{len(decoded), n})
			for ; n > 0; n-- {
				decoded = append(decoded, fileID)
			}
			fileID++
		} else {
			isFile = true
			if n > 0 {
				spaces = append(spaces, file{len(decoded), n})
			}
			for ; n > 0; n-- {
				decoded = append(decoded, -1)
			}
		}
	}
	return decoded, files, spaces
}

func run(decoded []int, files []file, spaces []file) []int {
	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		cur := decoded[file.i]
		for j := 0; j < len(spaces); j++ {
			spc := spaces[j]
			if spc.len >= file.len && spc.i < file.i {
				k := 0
				for k < file.len {
					decoded[spc.i] = cur
					decoded[file.i] = -1
					spc.i++
					file.i++
					k++
				}
				spc.len -= file.len
				spaces[j] = spc
				break
			}
		}
	}
  fmt.Println(decoded)
	return decoded
}
func part2(in string) int {
	dec := run(parse(in))
	res := 0
	for i, v := range dec {
		if v > 0 {
			res += i * v
		}
	}
	return res
}
