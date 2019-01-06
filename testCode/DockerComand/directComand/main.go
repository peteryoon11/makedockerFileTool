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

	//cmd := exec.Command("docker", "version")
	// docker run --rm -it ubuntu:16.04 /bin/bash

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

func StoreInfo() {

}
