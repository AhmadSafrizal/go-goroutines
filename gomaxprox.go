package goroutines

import (
	"testing"
	"fmt"
	"runtime"
)

func TestGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU = ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread = ", totalThread)
	
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Thread = ", totalGoroutine)
}