package range_demo_test

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	for i, c := range "This is golang" {
		fmt.Println(i, string(c))
	}

	nums := []int {1,2,3,4,5,6}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	nameAges := map[string]int {"freeman":35, "jenny":5}
	for k, v := range nameAges {
		fmt.Println(k, v)
	}
}
