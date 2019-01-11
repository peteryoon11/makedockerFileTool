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
		   docker 이미지 끼리만 연결 할때는(같은 서버의 컨테이너 끼리 연결 할때 )

		   docker run --link

		   -----------

		   1. mongodb 컨테이너 런
		   docker run --name db -d mongo

		   2. nginx 컨테이너를 생성 하면서 mongodb 이미지와 연결
		   docker run --name web  -d -p 80:80 --link db:db nginx

		   3. docker ps 명령어로 정보를 확인 할때 mongodb 에 names 라는 칸에  db, web/db  라고 떠있으면 된다.
		   docker ps

		   4. web container 안에서  아래의 명령어로 접속 가능
		   mongodb://db:27017/exapledb



		   다른 서버에 있는 컨테이너 끼리 연결 하려면 DockerFile 로 빌드 된 이미지가 필요하다.
		# 아래 를 실행 하기 전에 link 된 컨테이너가 실행 중이어야 한다.
		아래의 명령으로
		CMD env | grep _TCP = |\ sde

		docker run -d --link redis:redis --name redis_ambassador -p 6379:6379 svendowideit/ambassador
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
