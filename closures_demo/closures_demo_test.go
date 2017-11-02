package closures_demo_test

import (
	"fmt"
	"testing"
)

func initSeq() func() int {
	i := 0
	return func() int {
		i+=1
		return i
	}
}

func Test(t *testing.T) {
	seq := initSeq()
	fmt.Println(seq())
	fmt.Println(seq())
}