package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	/*
		Docker 베이스 이미지 생성하기

		sudo apt-get install debootstrap

		debootstrap trusty      trusty
		debootstrap <코드네임> <디렉토리>


		tar -C truly -c . | sudo docker import -trustly
	*/
	cmd := exec.Command("tr", "a-z", "A-Z")
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stdin = strings.NewReader("some input")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
