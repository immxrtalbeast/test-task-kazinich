package service

import (
	"math/rand/v2"
)

type RTPMultiplierService struct {
	targetRTP float64
}

func NewRTPMultiplierService(targetRTP float64) *RTPMultiplierService {
	return &RTPMultiplierService{
		targetRTP: targetRTP,
	}
}

func (s *RTPMultiplierService) Generate() float64 {
	ran := rand.Float64()
	if ran > s.targetRTP {
		return 1.0
	}
	return 10000.0
}
