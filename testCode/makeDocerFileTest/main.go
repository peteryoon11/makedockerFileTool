package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	/*
		need element
		from
		-> 어느 도커 이미지를 기반으로 만들건지
		maintainer
		-> 관리자
		add
		-> 현재 나의 공간에서 도커 이미지 내부로 추가할 파일들

	*/
	// sudo 를 붙여야 하는 경우에 대해서도 고려 해야 함
	//docker run --rm -it ubuntu:16.04 /bin/bash
	//docker run ubuntu:16.04
	temp := "docker build -t app "
	for num, item := range strings.Split(temp, " ") {
		fmt.Println(num, item)
	}
	exec.Command("docker")
}
func MakeDockerFile_cmd() {
	//temp := "from "
	//temp_Array := make([]string, 0)
	//	temp_Array = append(temp_Array, "from")
	//keyword := make([]string, 0)
	keyword := []string{"From", "Maintainer", "run", "volume", "workdir", "cmd", "expose"}
	os := []string{"ubuntu:latest"}
	/*
		sudo docker run --name hello-nginx -d -p 80:80 -v /root/data:/data hello:0.1

	*/
	for index, item := range keyword {
		fmt.Println(index, item)
	}
}
