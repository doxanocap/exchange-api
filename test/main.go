package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
   	var wg sync.WaitGroup
	ch := make(chan error, 1)
	go func() {
		wg.Add(1)
		fmt.Println("AUTH service shut down started...")
		for i := 0; true; i++ {
			if i == 2 {
				wg.Done()
				ch <- fmt.Errorf("AUTH service is not shutting down ERROR")
			}
			time.Sleep(1 * time.Second)
		}
	}()

	if err := <-ch; err != nil {
		fmt.Println(err.Error())
	}

    for i := 0; true; i++ {
        
            fmt.Println("waiting")
            time.Sleep(time.Second)
		}
    wg.Wait()
}
