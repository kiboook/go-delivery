package main

import (
	"go-delivery/config"
)

func main() {
	// 설정 파일 로드
	config.LoadConfig()

	// DB connect
	config.Connect()

}
