package main

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		c.JSON(http.StatusOK, gin.H{
			"version":    "1.0.0",
			"name":       "server1",
			"alloc":      bToMb(m.Alloc),
			"totalAlloc": bToMb(m.TotalAlloc),
			"sys":        bToMb(m.Sys),
			"numGC":      m.NumGC,
		})
	})

	r.Run()
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
