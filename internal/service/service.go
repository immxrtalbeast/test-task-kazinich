package service

import (
	"log/slog"
	"math/rand/v2"
)

type RTPMultiplierService struct {
	targetRTP float64
	log       *slog.Logger
}

func NewRTPMultiplierService(targetRTP float64, log *slog.Logger) *RTPMultiplierService {
	return &RTPMultiplierService{
		targetRTP: targetRTP,
		log:       log,
	}
}

func (s *RTPMultiplierService) Generate() float64 {
	ran := rand.Float64()

	if ran > s.targetRTP {
		return 1.0
	}
	return 10000.0

	// randomFactor := 0.0
	// low := ideal * (1.0 - randomFactor)
	// high := ideal * (1.0 + randomFactor)
	// multiplier := low + rand.Float64()*(high-low)
	// if multiplier > 10000.0 {
	// 	multiplier = 10000.0
	// }
	// if multiplier < 1.0 {
	// 	multiplier = 1.0
	// }
	// rounded := math.Round(multiplier*10) / 10 //Округление если надо
	// log.Info("Multiplier", slog.Float64("result", rounded))
	// return rounded
}
