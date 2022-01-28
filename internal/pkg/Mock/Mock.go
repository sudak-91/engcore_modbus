package Mock

import "math/rand"

func GenerateCoilMoсk(coilscount int) []byte {
	resultslice := make([]byte, coilscount)

	for k, _ := range resultslice {
		resultslice[k] = byte(rand.Intn(2))
	}
	return resultslice

}
