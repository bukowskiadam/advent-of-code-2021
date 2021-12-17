package utils

import (
	"fmt"
	"time"
)

func MeasureExecutionTime(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
