package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateUID() int64 {
	idString := fmt.Sprintf("%v%v", time.Now().UnixMilli(), rangeIn(1000, 9000))
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		panic(err)
	}
	return id
}

func rangeIn(low, hi int) int {
	rand.Seed(time.Now().UnixNano())
	return low + rand.Intn(hi-low)
}
