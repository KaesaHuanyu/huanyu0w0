package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println()
}

func hammingDistance(x int, y int) int {
	var hd int
	a := binary(x)
	b := binary(y)
	if x > y {
		t := len(a) - len(b)
		for i := 0; i < t; i++ {
			b = append(b, 0)
		}
		for i, v := range a {
			if v != b[i] {
				hd++
			}
		}
	} else if y > x {
		t := len(b) - len(a)
		for i := 0; i < t; i++ {
			a = append(a, 0)
		}
		for i, v := range b {
			if v != a[i] {
				hd++
			}
		}
	} else {
		return 0
	}

	return hd
}

func binary(x int) (y []int) {
	if x == 0 {
		return
	} else if x == 1 {
		y = append(y, 1)
		x = 0
		return
	}
	y = append(y, x%2)
	yy := binary(x / 2)
	y = append(y, yy...)
	return
}

func transNumber(l *ListNode, y []int) []int {
	y = append(y, l.Val)
	if l.Next != nil {
		y = transNumber(l.Next, y)
	}
	return y
}

func transInt(y []int) (x int) {
	for i := 0; i < len(y); i++ {
		n := y[i]
		for m := 0; m < len(y)-i-1; m++ {
			n *= 10
		}
		x += n
	}
	return
}

func transSlice(x int) (y []int) {
	if x == 0 {
		return []int{0}
	}
	t := x
	for t > 0 {
		y = append(y, 0)
		t /= 10
	}
	for i := 0; i < len(y); i++ {
		y[len(y)-i-1] = x % 10
		x /= 10
	}
	return
}

func transListNode(y []int) (l *ListNode) {
	if len(y) == 0 {
		return
	}
	l = &ListNode{
		Val:  y[0],
		Next: transListNode(y[1:]),
	}
	return
}
