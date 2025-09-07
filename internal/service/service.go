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
	const op = "service.GenerateMultiplier"
	s.log.With(
		slog.String("op", op),
		slog.Float64("rtp", s.targetRTP),
	)
	ideal := math.Sqrt(s.targetRTP) * 10000.0
	randomFactor := 0.3
	// Этот множитель можно регулировать чтобы брать бОльший диапазон мультипликаторов.
	// Погрешность при  randomFarctor = 0.3 и rtp = 0.5 бывает < 0.02, а разница между minMult и maxMult ~ 4000

	low := ideal * (1.0 - randomFactor)
	high := ideal * (1.0 + randomFactor)
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
