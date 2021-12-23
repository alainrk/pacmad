package main

import "math/rand"

func RandIntInRange(min, max int) int {
	return min + rand.Intn(max-min)
}
