package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var Colors = []string{
	"182,70,160",
	"12,199,233",
	"42,226,188",
	"177,231,77",
	"162,27,6",
	"66,181,209",
	"174,172,36",
	"31,29,15",
	"55,69,116",
}

func NewColorRGB() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%d, %d, %d", rand.Intn(255), rand.Intn(255), rand.Intn(255))
}
