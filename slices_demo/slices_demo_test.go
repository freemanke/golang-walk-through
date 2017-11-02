package slices_demo_test

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	s := make([]string, 3)
	s[0] = "0"
	s[1] = "1"
	s[2] = "2"
	s = append(s, "3")
	s = append(s, "4")
	s = append(s, "5")
	fmt.Println("Should be [0 1 2 3 4 5] ", s)
	fmt.Println("Should be [1 2]", s[1:3])
	
}