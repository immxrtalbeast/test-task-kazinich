package service

import (
	"log/slog"
	"math"
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

func (s *RTPMultiplierService) GenerateMultiplier() float64 {
	const op = "service.GenerateMultiplier3"
	s.log.With(
		slog.String("op", op),
		slog.Float64("rtp", s.targetRTP),
	)
	ideal := math.Sqrt(s.targetRTP) * 10000.0
	low := ideal * 0.7  // Эти множители можно регулировать чтобы брать бОльший диапазон мультипликаторов
	high := ideal * 1.3 // Погрешность при этих значениях < 0.02, а разница между minMult и maxMult ~ 4000
	multiplier := low + rand.Float64()*(high-low)
	if multiplier > 10000.0 {
		multiplier = 10000.0
	}
	if multiplier < 1.0 {
		multiplier = 1.0
	}
	rounded := math.Round(multiplier*10) / 10 //Округление если надо
	s.log.Info("Multiplier", slog.Float64("result", rounded))
	return rounded
}
