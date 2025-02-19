package metrics

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	messageCount int
	totalLatency int64
	mutex        sync.Mutex
)

// Latency 업데이트
func UpdateMetrics(latency int64) {
	mutex.Lock()
	messageCount++
	totalLatency += latency
	mutex.Unlock()
}

// 초당 Throughput 및 CPU, RAM 측정
func MonitorPerformance() {
	for {
		time.Sleep(1 * time.Second)
		mutex.Lock()

		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		avgLatency := int64(0)
		if messageCount > 0 {
			avgLatency = totalLatency / int64(messageCount)
		}

		fmt.Printf("[INFO] TPS: %d, Avg Latency: %d ms, CPU: %d cores, RAM: %.2fMB\n",
			messageCount, avgLatency, runtime.NumCPU(), float64(memStats.Alloc)/1024/1024)

		messageCount = 0
		totalLatency = 0
		mutex.Unlock()
	}
}
