package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(10 * time.Millisecond) // hapus ini bakal kelihatan hasilnya beda-beda
}

func DispalyNumber(number int) {
	fmt.Println("Dispaly", number)
}

func TestManyGoroutine(t *testing.T) {

	for i := 0; i <= 100000; i++ {
		go DispalyNumber(i)
	}

	time.Sleep(10 * time.Second)
}
