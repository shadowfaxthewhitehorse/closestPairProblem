package main

import (
    "fmt"
    "math"
)

type Point struct {
    x, y float64
}

func main() {
    points := []Point{
        {1, 2},
        {3, 4},
        {5, 6},
        {7, 8},
        {9, 10},
    }

    closestPair := closest(points)

    fmt.Printf("Closest pair of points: (%.2f, %.2f) and (%.2f, %.2f)\n", closestPair[0].x, closestPair[0].y, closestPair[1].x, closestPair[1].y)
}

func closest(points []Point) [2]Point {
    n := len(points)
    if n < 2 {
        return [2]Point{}
    }

    if n == 2 {
        return [2]Point{points[0], points[1]}
    }

    // Sort points by x-coordinate
    sortedPoints := make([]Point, n)
    copy(sortedPoints, points)
    quickSort(sortedPoints, 0, n-1)

    // Find the closest pair recursively
    return closestHelper(sortedPoints, 0, n-1)
}

func closestHelper(points []Point, left, right int) [2]Point {
    if left == right {
        return [2]Point{}
    }

    if right-left == 1 {
        return [2]Point{points[left], points[right]}
    }

    mid := (left + right) / 2
    leftClosest := closestHelper(points, left, mid)
    rightClosest := closestHelper(points, mid+1, right)
    closest := minDistance(leftClosest, rightClosest)

    strip := make([]Point, 0)
    for i := left; i <= right; i++ {
        if math.Abs(points[i].x-points[mid].x) < closest {
            strip = append(strip, points[i])
        }
    }

    stripClosest := stripClosest(strip, closest)
    return minDistance(closest, stripClosest)
}

func stripClosest(strip []Point, closest float64) [2]Point {
    n := len(strip)
    if n < 2 {
        return [2]Point{}
    }

    minDist := closest
    var closestPair [2]Point

    for i := 0; i < n-1; i++ {
        for j := i+1; j < n; j++ {
            dist := distance(strip[i], strip[j])
            if dist < minDist {
                minDist = dist
                closestPair = [2]Point{strip[i], strip[j]}
            }
        }
    }

    return closestPair
}

func minDistance(p1, p2 [2]Point) [2]Point {
    if p1 == [2]Point{} {
        return p2
    }
    if p2 == [2]Point{} {
        return p1
    }

    dist1 := distance(p1[0], p1[1])
    dist2 := distance(p2[0], p2[1])

    if dist1 < dist2 {
        return p1
    }
    return p2
}

func distance(p1, p2 Point) float64 {
    return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func quickSort(points []Point, left, right int) {
    if left < right {
        pivot := partition(points, left, right)
        quickSort(points, left, pivot-1)
        quickSort(points, pivot+1, right)
    }
}

func partition(points []Point, left, right int) int {
    pivotIndex := (left + right) / 2
    pivotValue := points[pivotIndex].x
    points[pivotIndex], points[right] = points[right], points[pivotIndex]

    storeIndex := left
    for i := left; i < right; i++ {
        if points[i].x < pivotValue {
            points[i], points[storeIndex] = points[storeIndex], points[i]
            storeIndex++
        }
    }

    points[storeIndex], points[right] = points[right], points[storeIndex]
    return storeIndex
}
