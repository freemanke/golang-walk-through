package interface_demo

import (
	"reflect"
	"math"
	"testing"
)

type IArea interface {
	Area() float64
}

type Circle struct {
	R float64
}

func (r Circle) Area() float64 {
	return math.Pi * math.Pow(r.R, 2)
}

type Rectange struct {
	Width  float64
	Height float64
}

func (r Rectange) Area() float64 {
	return r.Height * r.Width
}

func CaculateArea(i IArea) float64 {
	return i.Area()
}

func NullInterface(i interface {}) reflect.Value {
	return reflect.ValueOf(i)
}

func TestNullInterface(t *testing.T) {
	circle := Circle{R: 10}
	value := NullInterface(circle)
	t.Log(value)
}

func TestCaculateArea(t *testing.T) {
	circle := Circle{R: 10}
	area := CaculateArea(circle)
	if area != math.Pi*100 {
		t.Error()
	}
}

func TestMethodUsage(t *testing.T) {
	circle := Circle{R: 10}
	if circle.Area() != 100*math.Pi {
		t.Error("")
	}
}
