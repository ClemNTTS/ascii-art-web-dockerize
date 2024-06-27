package lib

import "fmt"

func Print(a, b, c, d, e, f, g, h string) {
	if len(a) > 0 {
		fmt.Println(a)
	}
	if len(b) > 0 {
		fmt.Println(b)
	}
	if len(c) > 0 {
		fmt.Println(c)
	}
	if len(d) > 0 {
		fmt.Println(d)
	}
	if len(e) > 0 {
		fmt.Println(e)
	}
	if len(f) > 0 {
		fmt.Println(f)
	}
	if len(g) > 0 {
		fmt.Println(g)
	}
	if len(h) > 0 {
		fmt.Println(h)
	}
	if a == "" && b == "" && c == "" && d == "" && e == "" && f == "" && g == "" && h == "" {
		fmt.Println((""))
	}
}
