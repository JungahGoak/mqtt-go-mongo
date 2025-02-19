package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type MqttData struct {
	ID        string    `json:"id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
}

func StartMessageHandler(messageCh <-chan []byte, dbClient *mongo.Client) {
	var wg sync.WaitGroup
	workerCount := 5 // 고루틴 수

	// Worker Pool (병렬화)
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			processMessages(workerID, messageCh, dbClient)
		}(i)
	}
	wg.Wait()
}

func processMessages(workerID int, messageCh <-chan []byte, dbClient *mongo.Client) {
	batch := make([]interface{}, 0, 100) // ✅ Batch Insert를 위한 버퍼
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case msg, ok := <-messageCh:
			if !ok {
				return
			}
			var data MqttData
			if err := json.Unmarshal(msg, &data); err != nil {
				log.Printf("❌ JSON 파싱 오류: %v", err)
				continue
			}
			batch = append(batch, data) // ✅ 배치에 데이터 추가

		case <-ticker.C:
			if len(batch) > 0 {
				saveToMongo(batch, dbClient)
				batch = batch[:0] // ✅ 배치 초기화
			}
		}
	}
}

func saveToMongo(batch []interface{}, dbClient *mongo.Client) {
	collection := dbClient.Database("robot_db").Collection("positions")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertMany(ctx, batch) // ✅ Batch Insert
	if err != nil {
		log.Printf("❌ MongoDB 저장 실패: %v", err)
	} else {
		fmt.Printf("✅ %d개의 데이터 저장 완료\n", len(batch))
	}
}
