package main

import (
	"encoding/json"
	"net/http"
	"runtime"
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
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
