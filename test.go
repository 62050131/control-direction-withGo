package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// x := 0
	// y := 0
	// current := 0
	// size := ""
	// space := 0

	ReadFile()

}

func ReadFile() {
	// open file
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	//  close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)
	text := ""
	size := ""
	arrText := make([]string, 100)
	for i := 0; scanner.Scan(); i++ {

		// get size of space map from the first line
		if i == 0 {
			size = scanner.Text()

		} else {
			//get direction from others line
			text = scanner.Text()
			arrText[i] = text
		}

		//make movement from direction
		Moving(arrText, StringToInt(size))
		//fmt.Printf("line %d: %s %s\n", i, text, size)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(StringToInt(size))

}

func StringToInt(size string) int {
	//convert string to int
	space, _ := strconv.Atoi(size)

	return space
}

//define movement from direction
func Moving(arrText []string, space int) (int, int, int) {

	turn := "N"
	arrDirect := [4]string{"N", "E", "S", "W"}
	current := 0
	x := 0
	y := 0
	for i := 0; i < len(arrText); i++ {
		//if turn left current will decrease and arrDirect value is go on the left side
		if arrText[i] == "L" {
			current -= 1
			if current < 0 {
				current = len(arrDirect) - 1
			}
			turn = arrDirect[current]

			//if turn right current will increase and arrDirect value is go on the right side
		} else if arrText[i] == "R" {
			current += 1
			if current > len(arrDirect)-1 {
				current = 0
			}
			turn = arrDirect[current]

			//change x or y value when forward
		} else if arrText[i] == "F" {

			if turn == "N" {
				y += 1
				y = pointY(y, space)

			} else if turn == "S" {
				y -= 1
				y = pointY(y, space)

			} else if turn == "E" {
				x += 1
				x = pointX(x, space)

			} else if turn == "w" {
				x -= 1
				x = pointX(x, space)

			}

		}
	}
	fmt.Printf(" %s:%d,%d\n", turn, x, y)
	return x, y, current
}

//limit x within space map
func pointX(x int, space int) int {
	X := 0
	if x < 0 {
		X = 0
	} else if x > space {
		X = space
	} else {
		X = x
	}
	return X
}

//limit y within space map
func pointY(y int, space int) int {
	Y := 0
	if y < 0 {
		Y = 0
	} else if y > space {
		Y = space
	} else {
		Y = y
	}
	return Y
}
