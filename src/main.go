package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	/*
		docker build <옵션> <Dockerfile 경로>

	*/

	cmd := exec.Command("docker", "version")
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
