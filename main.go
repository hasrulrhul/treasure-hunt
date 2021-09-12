package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// var r *rand.Rand
var matrix [6][8]string

// func init() {
// 	r = rand.New(rand.NewSource(time.Now().UnixNano()))
// }

var (
	// location start player
	x int = 4
	y int = 1
	// location treasure
	x1 int = 1
	y1 int = 1
)

func main() {
	// inisialisasi track
	MapTrack(x, y)
	// print track
	MapTrackPrint(x, y)
	// description option
	Description()

	reader := bufio.NewReader(os.Stdin)
	for {

		// CheckTrack(x, y)
		fmt.Print("selected track : ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if x == 4 && y == 6 {
			fmt.Println("track not available.")
			return
		}
		if strings.Compare("1", text) == 0 {
			obstacles := CheckObstacle(x-1, y)
			if obstacles != "#" {
				matrix[x][y] = "."
				x = UpTrack(x)
				if x == x1 && y == y1 {
					matrix[x][y] = "$"
					fmt.Println("treasure found.")
					MapTrackPrint(x, y)
					return
				}
				matrix[x][y] = "X"
				MapTrackPrint(x, y)
			}

		} else if strings.Compare("2", text) == 0 {
			obstacles := CheckObstacle(x+1, y)
			if obstacles != "#" {
				matrix[x][y] = "."
				x = DownTrack(x)
				if x+1 == x1 && y == y1 {
					matrix[x+1][y] = "$"
					fmt.Println("treasure found.")
					MapTrackPrint(x, y)
					return
				}
				matrix[x][y] = "X"
				MapTrackPrint(x, y)
			}

		} else if strings.Compare("3", text) == 0 {
			obstacles := CheckObstacle(x, y+1)
			if obstacles != "#" {
				matrix[x][y] = "."
				y = RightTrack(y)
				if x == x1 && y+1 == y1 {
					matrix[x][y+1] = "$"
					fmt.Println("treasure found.")
					MapTrackPrint(x, y)
					return
				}
				matrix[x][y] = "X"
				MapTrackPrint(x, y)
				if x == 4 && y == 6 {
					fmt.Println("track not available.")
					return
				}
			}
		} else {
			MapTrackPrint(x, y)
			fmt.Println("track not available.")
			return
		}
	}
}

// create track
func MapTrack(x, y int) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 8; j++ {
			if i == 0 || i == 5 || j == 0 || j == 7 {
				matrix[i][j] = "#"
			} else {
				if i == 2 && (j == 2 || j == 3 || j == 4) {
					matrix[i][j] = "#"
				} else if i == 3 && (j == 4 || j == 6) {
					matrix[i][j] = "#"
				} else if i == 4 && j == 2 {
					matrix[i][j] = "#"
				} else {
					if i == x && j == y {
						matrix[i][j] = "X"
					} else {
						matrix[i][j] = "."
					}
				}
			}
		}
	}
}

// print map track
func MapTrackPrint(x, y int) {
	fmt.Println("Map Track Treasure Hunt")
	fmt.Println()
	for i := 0; i < 6; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
	if matrix[x][y] != "$" {
		CheckTrack(x, y)
	}
	fmt.Println()
}

func UpTrack(x int) int {
	return x - 1
}

func DownTrack(x int) int {
	return x + 1
}

func RightTrack(y int) int {
	return y + 1
}

func CheckObstacle(x, y int) string {
	if matrix[x][y] == "#" {
		fmt.Println("there are obstacles")
		return "#"
	}
	return ""
}

func Description() {
	fmt.Println()
	fmt.Println(".:: Treasure Hunt ::.")
	fmt.Println("---------------------")
	fmt.Println("Description")
	fmt.Println("# : obstacle")
	fmt.Println("X : start position ")
	fmt.Println("$ : location treasure ")
	fmt.Println("----------------------")
	fmt.Println()
	MapTrackPrint(x, y)
}

func OptionTrack() {
	fmt.Println("Option Track")
	fmt.Println("------------")
	fmt.Println("1: up")
	fmt.Println("2: down")
	fmt.Println("3: right")
	fmt.Println("---------")
}

func CheckTrack(x, y int) {
	fmt.Println("Option Track")
	fmt.Println("------------")
	if matrix[x][y+1] != "$" {
		// check track up x-1
		if matrix[x-1][y] != "#" {
			fmt.Println("1: up")
		}
		// check track down x+1
		if matrix[x+1][y] != "#" {
			fmt.Println("2: down")
		}
		// check track right y+1
		if matrix[x][y+1] != "#" {
			fmt.Println("3: right")
		}
	}
}
