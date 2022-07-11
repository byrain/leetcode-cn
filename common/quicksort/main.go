package main

import "fmt"

func quickSort(s []int, l, h int) {
	if l >= h {
		return
	}
	tempL := l
	tempH := h

	for {
		for s[tempH] > s[l] {
			tempH--
		}
		for s[tempL] <= s[l] && tempL < tempH {
			tempL++
		}
		if tempL >= tempH {
			break
		}
		s[tempL], s[tempH] = s[tempH], s[tempL]
	}
	s[tempL], s[l] = s[l], s[tempL]
	quickSort(s, l, tempL-1)
	quickSort(s, tempL+1, h)
}

func main() {
	s := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	quickSort(s, 0, len(s)-1)
	fmt.Println("--------1-----------")
	fmt.Printf("%#+v\n", s)
	fmt.Println("--------2-----------")

}
