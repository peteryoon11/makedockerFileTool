package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	tempLocate := "temp"
	// 이 부분을 절대 경로로 받아올까...
	/* testUseFulTool 에서 pwd 명령을 날려서
	경로를 가져오는게 있는데 이거는 윈도우랑
	윈도우는 chdir 를 사용
	*/
	DockerFileLocate := "./" + tempLocate + "/DockerFile"
	// logger 의 위치 에 대한 테스트 중
	// ./ ../
	fpLog, err := os.OpenFile(DockerFileLocate, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		//panic(err)
		//fmt.Println(err)
		log.Println(err)
		/* 	workerRecorder = log.New(fpLog, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
		FileControll.CheckOSAndMakeFile(loggerLocateWithFile, workerRecorder)
		// 로거가 생성 되기 전에 폴더가 없으면 폴더를 생성하고 나서 다시 한번 파일을 열어서 로그 저장을 함
		fpLog, err = os.OpenFile(loggerLocateWithFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			workerRecorder.Panicln(err)
		} */

	}

	//multiWriter := io.MultiWriter(fpLog, os.Stdout)
	//log.SetOutput(multiWriter)
	//fpLog
	//workerRecorder := log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	if _, err := fpLog.Write([]byte("appended some data\n")); err != nil {
		log.Fatal(err)
	}
	if err := fpLog.Close(); err != nil {
		log.Fatal(err)
	}
	//workerRecorder.SetOutput(multiWriter)
	fmt.Println("test")
}
