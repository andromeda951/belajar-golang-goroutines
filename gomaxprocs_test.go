package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetGomaxprocs(t *testing.T) {

	totalCpu := runtime.NumCPU()
	fmt.Println("CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Goroutine", totalGoroutine)
}
