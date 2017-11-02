package maps_demo_test

import (
	"testing"
)

func TestMaps(t *testing.T) {
	nameAges := make(map[string]int)
	nameAges["freeman"] = 35
	nameAges["jenny"] = 5

	for name := range nameAges {
		println(name, nameAges[name])
	}
}