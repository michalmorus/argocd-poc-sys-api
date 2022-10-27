package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		data := map[string]any{
			"version":    "1.0.0",
			"name":       "server1",
			"alloc":      bToMb(m.Alloc),
			"totalAlloc": bToMb(m.TotalAlloc),
			"sys":        bToMb(m.Sys),
			"numGC":      m.NumGC,
			"time":       time.Now().Format("15:04:05.000000"),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	})

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}

	fmt.Println("listen on ", port)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
