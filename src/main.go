package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

/*
1. CI(JENKINS) Pipeline 구축
 => jenkins 에서의 ci 를 모델로 해당 부분을 구현

3. Docker 기반 개발환경(PaaS) 구축
=> Docker 파일 로 서버와 서버 내부의 패키 지 등을 설정 가능하게 해서 자동으로 (AWS에서의 웹 콘솔처럼) 원하는 개발 환경을 구축하는 부분을 자동화 한다.

*/
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
