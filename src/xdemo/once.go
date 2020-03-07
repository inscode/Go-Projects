package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	o := &sync.Once{}
	go Do(o)
	go Do(o)
	time.Sleep(time.Second)
}

func Do(o *sync.Once) {
	fmt.Println("start do")

	o.Do(func() {
		fmt.Println("doing something")
	})

	fmt.Println("do end")
}
