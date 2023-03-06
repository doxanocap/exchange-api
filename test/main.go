package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)
	
func main() {
	ticker := time.NewTicker(1 * time.Second)
	err := make(chan error, 1)

    fmt.Println(time.Now().Unix())

	fmt.Println("start")
    var wg sync.WaitGroup
    wg.Add(1)

    go func() {
        for {
            select {
            case err1 := <-err:
                fmt.Println(err1.Error())

            case <-ticker.C:
                fmt.Println("tick")
                err1 := ToDo()
                err <- err1
                err <- err1
                err <- err1

            }
        }
    }()

    for {

    }
}

func ToDo() error {
	return errors.New("error in todo")
}