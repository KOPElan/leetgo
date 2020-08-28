package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	fmt.Println("ready to go...")
	str, err := ioutil.ReadFile("case.txt")
	if err != nil {
		panic(err)
	}

	minNumberOfFrogs(string(str))
}

func minNumberOfFrogs(croakOfFrogs string) int {
	start := time.Now()
	var x, y, c, r, o, a, k int
	for _, s := range croakOfFrogs {
		switch {
		case s == 'c':
			c++
			if y > 0 {
				y--
			} else {
				x++
			}
			break
		case r < c && s == 'r':
			r++
			break
		case o < r && s == 'o':
			o++
			break
		case a < o && s == 'a':
			a++
			break
		case k < a && s == 'k':
			k++
			y++
			break
		default:
			return -1
		}
	}
	if x != y {
		return -1
	}
	fmt.Println("[耗时]", time.Now().Sub(start))
	return x
}
