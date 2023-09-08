package utils

import (
	"log"
	"strconv"
)

func StringToInt(value string) int {

	v, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalln(err)
	}

	return v
}

func StringToBool(value string) bool {
	return value == "true"
}
