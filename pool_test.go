package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool sync.Pool

	pool.Put("Eko")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")

	// Using Pointer
	// firstName := "Eko"
	// middleName := "Kurniawan"
	// lastName := "Khannedy"

	// pool.Put(&firstName)
	// pool.Put(&middleName)
	// pool.Put(&lastName)

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second) // add this line
			pool.Put(data)
		}()
	}

	time.Sleep(5 * time.Second)
}
