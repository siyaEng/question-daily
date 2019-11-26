package main

import "fmt"

func main() {
	s := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := s[2:6]  // 2,3,4,5 len 4 cap
	s2 := s1[3:5] // 5,6 len 2

	fmt.Printf("s1=%v, len(s1)=%d cap(s1)=%d", s1, len(s1), cap(s1))
	fmt.Println()
	fmt.Printf("s2=%v, len(s2)=%d cap(s2)=%d", s2, len(s2), cap(s2))

}
