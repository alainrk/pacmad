package main

import "math/rand"

func RandIntInRange(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandFloatInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
