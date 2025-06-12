package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Audit struct {
	Method        string    `json:"method"`
	Path          string    `json:"path"`
	ContentLength int64     `json:"content_length"`
	Ts            time.Time `json:"ts"`
}

var auditLog = make(chan Audit)

func init() {
	go func() {
		file, err := os.OpenFile("audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed to open audit.log: %v", err)
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
		encoder := json.NewEncoder(file)
		for audit := range auditLog {
			if err := encoder.Encode(audit); err != nil {
				log.Printf("failed to write audit log: %v", err)
			}
		}
	}()
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Method: %s Path: %s Duration: %s Size: %d B\n", r.Method, r.URL.Path, duration, r.ContentLength)

		auditLog <- Audit{
			Method:        r.Method,
			Path:          r.URL.Path,
			ContentLength: r.ContentLength,
			Ts:            start,
		}
	})
}
