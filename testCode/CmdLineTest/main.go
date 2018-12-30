package main

import (
	"fmt"
	"strings"
)

func main() {
	/* temp := make([]string, 0)
	fmt.Println(temp) */
	testString := "docker run --rm -it ubuntu:16.04 /bin/bash"
	for num, item := range strings.Split(testString, " ") {
		fmt.Println(num, item)
	}
}
