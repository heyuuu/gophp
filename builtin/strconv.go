package builtin

import (
	"log"
	"strconv"
)

func Atoi(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
