package mqtt

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewClient(messageCh chan<- []byte) mqtt.Client {
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://mosquitto:1883").             // ✅ MQTT 브로커 주소
		SetClientID("go_mqtt_subscriber").             // ✅ Client ID
		SetProtocolVersion(4).  		       // MQTT 3.1.1 (v4) 사용 설정
		SetCleanSession(false).                        // ✅ 이전 세션 유지 (장애 발생 후 메시지 다시 받음)
		SetAutoReconnect(true).                        // ✅ 자동 재연결
		SetKeepAlive(60 * time.Second).                // ✅ 60초마다 Ping 전송
		SetConnectRetry(true).                         // ✅ 재연결 시도 활성화
		SetOnConnectHandler(func(client mqtt.Client) { // ✅ MQTT 브로커 연결 시 콜백함수
			log.Printf("[MQTT] ✅ MQTT 연결 성공!")
			// ✅ MQTT 브로커 연결 성공 시 'robot/position' topic subscribe
			if token := client.Subscribe("robot/position", 1, func(client mqtt.Client, msg mqtt.Message) {
				messageCh <- msg.Payload() // ✅ 채널로 메시지 전송
			}); token.Wait() && token.Error() != nil {
				log.Fatalf("[MQTT] ❌ 구독 실패: %v", token.Error())
			}
		}).
		SetConnectionLostHandler(func(client mqtt.Client, err error) {
			log.Printf("[MQTT] ⚠️ MQTT 연결 끊김 : %v\n", err)
		})

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return client
}
