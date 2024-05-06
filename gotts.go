package main

import (
	"fmt"
	"sync"

	"gotts/tts"

)

func main() {
	fmt.Print("Hello\n")



	var wg sync.WaitGroup
	
	wg.Add(1)
	go tts.StartRESTInterface()

	wg.Wait()
}