package helper

import (
	"math/rand"
	"strconv"
)

func GenerateBureauReference() string {

	random := rand.Intn(900000) + 100000
	return "BURE" + strconv.Itoa(random)
}

func GenrateUserReference() string {
	random := rand.Intn(900000) + 100000
	return "USER" + strconv.Itoa(random)
}

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
