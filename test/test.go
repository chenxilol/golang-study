package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{
		"sdfs", "sdfsag",
	}
	s1 := strings.Join(s, "")
	fmt.Println(s1)
}
