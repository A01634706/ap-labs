// Copyright ©️ 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 156.

// Package geometry defines simple types for plane geometry.
//!+point
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

type Point struct{ x, y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p Point) X() float64{
	return p.x
}

func (p Point) Y() float64{
	return p.y
}
//Retrieved from https://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/
func onSegment(p, q, r Point) bool{
	if(q.X() <= math.Max(p.X(), r.X()) && q.X() >= math.Min(p.X(), r.X()) &&
	q.y <= math.Max(p.y, r.y) && q.y >= math.Min(p.y, r.y)){
		return true
	}
	return false
}

func orientation(p, q, r Point) int{
	val := (q.Y() - p.Y()) * (r.X() - q.X()) - (q.X() - p.X()) * (r.Y() - q.Y())

	if (val == 0){
		return 0
	}

	if (val > 0){
		return 1
	}else {
		return 2
	}
}

func doIntersect(p1, q1, p2, q2 Point) bool {
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	if (o1 != o2 && o3 != o4){
		return true
	}

	if (o1 == 0 && onSegment(p1, p2, q1)){
		return true
	}

	if (o2 == 0 && onSegment(p1, q2, q1)){
		return true
	}

	if (o3 == 0 && onSegment(p2, p1, q2)){
		return true
	}

	if (o4 == 0 && onSegment(p2, q1, q2)){
		return true
	}

	return false
}

//!-point

//!+path

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func main() {
	if(len(os.Args) < 2){
		fmt.Println("Error. Faltan argumentos")
		return
	}
	var arr Path;
	sides, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < sides; i++ {
		p := Point{(rand.Float64()*-100) + 100, (rand.Float64()*-100) + 100}
		if (i==sides-1) {
			for j := 0; j<10; j++{
				if (doIntersect(arr[len(arr)-2],arr[len(arr)-1],arr[len(arr)-1],p)){
					p = Point{(rand.Float64()*-100) + 100, (rand.Float64()*-100) + 100}
				}else{
					break
				}
			}
		} 
		arr = append(arr, p)
	}

	arr = append(arr, arr[0])

	fmt.Println("Perimetro: ")
	fmt.Println(arr.Distance())
}

//!-path