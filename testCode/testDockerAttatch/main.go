package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

/*

test docker attatch

컨테이너 띄워두고 명령 추가 하기

*/
func main() {
	cmd := exec.Command("docker", "run", "ubuntu:16.04")
	/* cmd.Stdin = strings.NewReader("some input")
	 */
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out.String())
}
