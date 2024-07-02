package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	a float64
	b float64
}

func (s *Rectangle) area() float64 {
	return s.a * s.b
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (s *Triangle) area() float64 {
	p := (s.a + s.b + s.c) / 2
	return math.Sqrt(p * (p - s.a) * (p - s.b) * (p - s.c))
}

type Line struct {
	a float64
}

func (s *Line) SetLine(newA float64) {
	s.a = newA
}

func main() {
	circle1, rectangle1, triangle1, line1 := &Circle{6}, &Rectangle{5, 12}, &Triangle{3, 4, 5}, &Line{7}

	c, err := calculateArea(circle1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Круг: радиус %g\nПлощадь: %v\n", circle1.r, c)
	}

	r, err := calculateArea(rectangle1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Прямоугольник: ширина %g, высота %g\nПлощадь: %g\n", rectangle1.a, rectangle1.b, r)
	}

	t, err := calculateArea(triangle1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Треугольник: стороны a: %g, b: %g, c: %g,\nПлощадь: %g\n", triangle1.a, triangle1.b, triangle1.c, t)
	}

	l, err := calculateArea(line1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Линия: длина %g\nПлощадь: %g\n", line1.a, l)
	}
}

func calculateArea(s any) (float64, error) {
	e, ok := s.(Shape)
	if !ok {
		return 0, fmt.Errorf("ошибка: переданный объект не является фигурой")
	}
	return e.area(), nil
}
