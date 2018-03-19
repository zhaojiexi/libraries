package test

import "math/rand"

func RandomInt ()int{

	return rand.Intn(10)
}

func Same ()bool{
	return  rand.Intn(10)<10
}