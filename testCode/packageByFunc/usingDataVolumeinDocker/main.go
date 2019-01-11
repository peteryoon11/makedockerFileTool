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
		Docker 데이터 볼륨 사용하기
		* Docker 데이터 볼륨은 데이터를 컨테이너가 아닌 호스트에 저장하는 방식
		* 컨테이너 끼리 데이터를 공유 할때 사용

		docker run -i -t --name hello-volume -v /data ubntu /bin/bash
		cd /data
		touch hello
		exit

		inspect 명령으로 hello-volume 컨테이너의 데이터 볼륨 경로 확인
		docker inspect -f "{{.Volumes}}" hello-volume

		ls 로 /var/lib/docker/vfs/dir/ 의 파일 목록을 출력

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
