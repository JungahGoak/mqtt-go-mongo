package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatalf("[MongoDB] ❌ MongoDB 클라이언트 생성 실패: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("[MongoDB] ❌ MongoDB 연결 실패: %v", err)
	}

	log.Println("[MongoDB] ✅ 연결 성공")
	return client
}
