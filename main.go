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
	"github.com/prometheus/common/log"
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
		log.Info("starting server", "port", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server failed to start %w", err)
			panic("fatal")
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	log.Info("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server forced to shutdown: %w", err)
		panic("fatal")
	}
	log.Info("server exiting")
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
