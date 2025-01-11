package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Response 标准响应结构
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// LoggingMiddleware 日志中间件
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	}
}

// HealthCheckHandler 健康检查
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  200,
		Message: "Service is healthy",
	}
	json.NewEncoder(w).Encode(response)
}

// HelloWorldHandler 主页处理
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  200,
		Message: "Hello, World!",
		Data:    map[string]string{"version": "1.0.0"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// 注册路由
	http.HandleFunc("/", LoggingMiddleware(helloWorldHandler))
	http.HandleFunc("/health", LoggingMiddleware(healthCheckHandler))

	// 启动服务器
	port := ":8089"
	fmt.Printf("Starting server on %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}