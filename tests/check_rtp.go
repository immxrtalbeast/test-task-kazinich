// Тестирую руками
package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"

	"github.com/immxrtalbeast/rtp-multiplier/internal/service"
)

func main() {
	targetRTP := 0.5
	if targetRTP <= 0.0 || targetRTP > 1.0 {
		panic("rtp should be ∈(0, 1.0]")
	}
	log := setupLogger()
	service := service.NewRTPMultiplierService(targetRTP, log)

	count := 10000
	zeros := 0
	minMultiplier := 10000.0
	maxMultiplier := 0.0
	randomFloats := make([]float64, count)
	seq1 := make([]float64, count)
	var sumRandom float64
	var sumSeq1 float64
	for i := 0; i < count; i++ {
		randomFloats[i] = 1 + rand.Float64()*9999
		sumRandom += randomFloats[i]
	}
	for i := 0; i < count; i++ {
		multiplier := service.Generate()
		if multiplier > maxMultiplier {
			maxMultiplier = multiplier
		}
		if multiplier < minMultiplier {
			minMultiplier = multiplier
		}
		if multiplier > randomFloats[i] {
			seq1[i] = randomFloats[i]
			sumSeq1 += randomFloats[i]
		} else {
			seq1[i] = 0
			zeros += 1
		}

	}
	// fmt.Println("RandomFloats \n", randomFloats)
	// fmt.Println("Seq1 \n", seq1)
	calculatedRTP := sumSeq1 / sumRandom
	fmt.Println("Zeros: \n", zeros)
	fmt.Printf("minMult:%.1f\n", minMultiplier)
	fmt.Printf("maxMult:%.1f\n", maxMultiplier)
	fmt.Printf("Calculated RTP: %.4f\n", calculatedRTP)
	// fmt.Printf("Target RTP: %.4f\n", targetRTP)
	// fmt.Printf("Calculated RTP: %.4f\n", calculatedRTP)
	fmt.Printf("Difference: %.4f\n", calculatedRTP-targetRTP)
}
func setupLogger() *slog.Logger {
	var log *slog.Logger

	log = slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	return log
}
