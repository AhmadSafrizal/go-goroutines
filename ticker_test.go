package goroutines

import (
	"fmt"
	"time"
	"testing"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	
	go func() {
		time.Sleep(10 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}