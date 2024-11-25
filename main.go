package main

import (
	"GoWD/modules/hello"
	"GoWD/modules/tcp"
	"sync"
)

func StartServerRoutine(wg *sync.WaitGroup) {
	defer wg.Done()
	tcp.StartServer()
}

func StartClientRoutine(wg *sync.WaitGroup, address string) {
	defer wg.Done()
	tcp.StartClient(address)
}

func main() {
	hello.SayHello()
	var wg sync.WaitGroup

	wg.Add(2)

	go StartServerRoutine(&wg)
	go StartClientRoutine(&wg, "localhost:8080")

	wg.Wait()

}
