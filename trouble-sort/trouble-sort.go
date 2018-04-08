package main

import "fmt"

func main() {
	var numLines uint
	fmt.Scanf("%d", &numLines)
	for i := uint(1); i < numLines+1; i++ {
		var length uint
		var list []uint
		fmt.Scanf("%d", &length)
		if length == 0 {
			fmt.Scanf("%s")
			fmt.Printf("Case #%d: OK\n", i)
			continue
		}
		for j := uint(0); j < length; j++ {
			var val uint
			fmt.Scanf("%d", &val)
			list = append(list, val)
		}
		troubleSort(list)
		e := indexOfError(list)
		fmt.Printf("Case #%d: ", i)
		if e == -1 {
			fmt.Println("OK")
		} else {
			fmt.Println(e)
		}
	}
}

func indexOfError(list []uint) int {
	for i := range list[0 : len(list)-1] {
		if list[i] > list[i+1] {
			return i
		}
	}
	return -1
}

func troubleSort(list []uint) {
	done := false
	for !done {
		done = true
		for i := 0; i < len(list)-2; i++ {
			if list[i] > list[i+2] {
				done = false
				s := list[i : i+3]
				// https: //stackoverflow.com/a/19239850
				for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
					s[i], s[j] = s[j], s[i]
				}
			}
		}
	}
}
