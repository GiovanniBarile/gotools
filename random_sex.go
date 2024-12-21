package gotools

import "math/rand"

// Return a random sex between M and F
func GetRandomSex() string {
	return []string{"M", "F"}[rand.Intn(2)]
}
