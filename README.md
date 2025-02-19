# mqtt-go-mongo

---
## 파일구조

```
mqtt-go-mongo/                  
├── cmd/                        # 애플리케이션 실행 엔트리포인트
│   └── main.go                 # 메인 함수, 초기화 및 실행
├── internal/                   # 주요 내부 모듈 (외부 접근 불가)
│   ├── mqtt/                   # MQTT 관련 로직
│   │   └── client.go           # MQTT 클라이언트 설정 및 처리
│   ├── mongodb/                # MongoDB 관련 로직
│   │   └── client.go           # MongoDB 클라이언트 설정 및 처리
│   └── handlers/               # 데이터 핸들링 모듈
│       └── handler.go          # 데이터 처리 및 비즈니스 로직 연결
├── pkg/                        # 공용 패키지 디렉터리 (다른 모듈에서 사용 가능)
│   ├── config/                 # 환경 설정 로직
│   │   └── config.go           # 설정 파일
│   └── models/                 # 데이터 모델 정의
│       └── location.go         # 위치 데이터 구조체
├── Dockerfile                  
└── docker-compose.yml          

```