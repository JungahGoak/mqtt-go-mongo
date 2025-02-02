package main

import (
	"mqtt-go-mongo/internal/handlers"
	"mqtt-go-mongo/internal/mongodb"
	"mqtt-go-mongo/internal/mqtt"
)

func main() {
	messageCh := make(chan []byte, 1000) // ✅ 메시지 버퍼 채널 생성

	mqtt.NewClient(messageCh)                            // ✅ MQTT 클라이언트 초기화
	dbClient := mongodb.NewClient()                      // ✅ MongoDB 클라이언트 초기화
	go handlers.StartMessageHandler(messageCh, dbClient) // ✅ 메시지 처리 고루틴 시작

	select {} // 메인 루프 유지
}
