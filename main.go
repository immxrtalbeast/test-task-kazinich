package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/immxrtalbeast/rtp-multiplier/internal/controller"
	"github.com/immxrtalbeast/rtp-multiplier/internal/service"
)

func main() {
	rtp := fetchRTP()
	rtpService := service.NewRTPMultiplierService(rtp)
	rtpController := controller.NewRTPController(*rtpService)

	router := gin.Default()
	router.GET("/get", rtpController.GetMultiplier)

	srv := &http.Server{
		Addr:         ":64333",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("fatal")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic("fatal")
	}
}

func fetchRTP() float64 {
	var rtp float64

	flag.Float64Var(&rtp, "rtp", 0.9, "Target RTP value (0 < rtp <= 1.0)")
	flag.Parse()
	if rtp <= 0 || rtp > 1 {
		panic("rtp should be ∈(0, 1.0]")
	}

	return rtp
}
